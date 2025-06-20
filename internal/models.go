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

type ImgExt struct {
	Name   string
	Status bool
}

func GetExtensionsList(enebled []string) []ImgExt {
	var list []ImgExt
	var statusFlag = false

	if len(enebled) == 0 {
		statusFlag = true
	}

	for _, ext := range Extensions {
		list = append(list,
			ImgExt{
				Name:   ext,
				Status: statusFlag,
			},
		)
	}

	if len(enebled) > 0 {
		for _, ext := range enebled {
			for i := range list {
				if ext == list[i].Name {
					list[i].Status = true
				}
			}
		}
	}

	return list
}
