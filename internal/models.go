package internal

// Extensions accepted file extensions for metadata extraction
var Extensions = []string{
	".jpg", ".jpeg",
	".png", ".cr2",
	".tif", ".tiff",
	".avif", ".apng",
	// ".webp",
	// ".nef", ".bmp",
	// ".arw", ".gif",
	// ".svg", ".ico",
}

// FileData holds metadata information about a file.
type FileData struct {
	Path     string
	Date     string
	Device   string
	Location string
	Size     int64
}

// GetExtensionsMap returns a map of enabled file extensions for image processing.
func GetExtensionsMap(enabled []string) map[string]bool {
	result := make(map[string]bool)

	for _, v := range enabled {
		result[v] = true
	}

	if enabled == nil {
		for _, v := range Extensions {
			result[v] = true
		}
	}
	return result
}
