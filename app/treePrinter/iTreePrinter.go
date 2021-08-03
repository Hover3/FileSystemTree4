package treePrinter

import "FileSystemTree4.mod/FileSystemTree4/domain/folderTree"

type iTreePrinter interface {
	PrintRecurrent(f *folderTree.FolderInfo, prefix string, last bool)
	PrintFolders(f *folderTree.FolderInfo, prefix string, last bool)
	PrintFiles(f *folderTree.FolderInfo, prefix string)
}
