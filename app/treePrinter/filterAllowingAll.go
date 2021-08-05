package treePrinter

import "FileSystemTree4/domain/folderTree"

type FilterAllowingAll struct {

}

func (filter *FilterAllowingAll) IsFileMatchFilter (f *folderTree.FileInfo) bool {
	return true
}


