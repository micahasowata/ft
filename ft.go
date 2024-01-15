package ft

var fileExtensions = map[string]string{
	".mp3":  "Audio",
	".wav":  "Audio",
	".aac":  "Audio",
	".flac": "Audio",
	".m4a":  "Audio",
	".mp4":  "Video",
	".avi":  "Video",
	".mkv":  "Video",
	".mov":  "Video",
	".wmv":  "Video",
	".jpg":  "Images",
	".jpeg": "Images",
	".png":  "Images",
	".gif":  "Images",
	".bmp":  "Images",
	".tiff": "Images",
	".tif":  "Images",
	".pdf":  "Documents",
	".docx": "Documents",
	".xlsx": "Documents",
	".pptx": "Documents",
	".txt":  "Documents",
	".mobi": "Documents",
	".epub": "Documents",
	".zip":  "Archives",
	".rar":  "Archives",
	".7z":   "Archives",
	".tar":  "Archives",
	".gz":   "Archives",
	".exe":  "Executables",
	".app":  "Executables",
	".dmg":  "Executables",
	".bin":  "Executables",
	".msi":  "Executables",
	".deb":  "Executables",
	".csv":  "Data",
	".json": "Data",
	".xml":  "Data",
}

func GetCategory(extension string) (string, error) {
	category, found := fileExtensions[extension]
	if !found {
		return "", ErrNotFound
	}
	return category, nil
}
