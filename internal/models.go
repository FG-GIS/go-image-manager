package internal

var Extensions = []string{
	".jpg", ".jpeg",
	".png", ".gif",
	".webp", ".tif",
	".tiff", ".bmp",
	".svg", ".ico",
	".avif", ".apng",
	".cr2", ".nef",
	".arw",
}

type FileData struct {
	Path     string
	Date     string
	Device   string
	Location string
}

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
