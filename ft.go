package ft

/*
GetFileCatgory takes in a file extension, stating with the . e.g .png, .txt which can be gotten from
filepath.Ext()

It returns a string on success or ft.ErrNotFound when the extension can not be located on the
list of files
*/
func GetFileCategory(extension string) (string, error) {
	category, found := fileExtensions[extension]
	if !found {
		return "", ErrNotFound
	}
	return category, nil
}
