package treePrinter

import "FileSystemTree4/domain/folderTree"

type iTreePrinter interface {
	PrintRecurrent(f *folderTree.FolderInfo, prefix string, last bool)
	printFolderRow(f *folderTree.FolderInfo, prefix string, last bool)
	printFiles(f *folderTree.FolderInfo, prefix string)
}
