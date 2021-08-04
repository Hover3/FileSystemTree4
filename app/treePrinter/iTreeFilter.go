package treePrinter

import "FileSystemTree4.mod/FileSystemTree4/domain/folderTree"

type ITreeFilter interface {
	IsFileMatchFilter (f *folderTree.FileInfo) bool
}
