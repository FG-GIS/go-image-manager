package internal

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// FileScanner scans the specified directory for files, excluding directories listed in exclusions.
// It returns a slice of file paths found in the directory.
// If an error occurs during the scan, it prints an error message and returns nil.
// Parameters:
//
//	root       - the root directory to scan.
//	exclusions - a slice of directory names to exclude from the scan.
//
// Returns:
//
//	A slice of file paths found in the directory, or nil if an error occurs.
func FileScanner(root string, exclusions []string, verbose bool) []string {
	var resultList []string

	err := filepath.WalkDir(root,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				if os.IsPermission(err) {
					if verbose {
						fmt.Printf("Error accessing %q: %v\n", path, err)
					}
					return filepath.SkipDir
				} else {
					return err
				}
			}
			name := strings.ToLower(d.Name())
			if d.IsDir() && slices.Contains(exclusions, name) {
				if verbose {
					fmt.Printf("Skipping directory %q due to exclusion\n", path)
				}
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

// FilterImages filters a slice of image file paths based on allowed file extensions provided in extMap.
// For each file with an allowed extension, it retrieves the file size and returns a slice of FileData
// containing the file path and size. If the file cannot be stat-ed, the size is set to zero.
//
// Parameters:
//
//	imgSlice - a slice of image file paths to filter.
//	extMap   - a map of allowed file extensions (keys are extensions, values are booleans).
//
// Returns:
//
//	A slice of FileData structs for files with allowed extensions, including their paths and sizes.
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
