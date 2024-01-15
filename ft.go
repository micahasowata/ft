package ft

import "strings"

/*
GetFileCatgory takes in a file extension which must begin with the leading dot e.g .png, .txt. It returns a string on success or one of these two errors:

1. ErrInvalidFormat - if there is no leading dot in the extension string

2. ErrNotFound - if extension cannot be located on the list of files
*/
func GetFileCategory(extension string) (string, error) {
	if !strings.HasPrefix(extension, ".") {
		return "", ErrInvalidFormat
	}
	category, found := fileExtensions[extension]
	if !found {
		return "", ErrNotFound
	}
	return category, nil
}
