package treePrinter

import (
	"FileSystemTree4.mod/FileSystemTree4/domain/folderTree"
	"strings"
)

type FilterExtensions struct {
	AllowedExtensions []string
}

func (filter *FilterExtensions) IsFileMatchFilter (f *folderTree.FileInfo) bool {
	ext:= DetectFileExtension(f.FileName)

for i:=range filter.AllowedExtensions {
	if ext==filter.AllowedExtensions[i] {
		return true
	}
}

	return false
}
func DetectFileExtension(fileName string) string {

	lastDotPosition:=strings.LastIndex(fileName, ".")
	if lastDotPosition==-1 {
		return ""
	}
	extension:=""
	for i, ch:=range fileName {
		if i>=lastDotPosition {
			extension=extension + string(ch)
		}

	}

	return extension
}