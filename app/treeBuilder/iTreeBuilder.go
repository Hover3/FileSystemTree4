package treeBuilder

import "FileSystemTree4.mod/FileSystemTree4/domain/folderTree"

type ITreeBuilder interface {
	ScanRecurrent(f *folderTree.TreeFolderInfo)
	Scan(f *folderTree.TreeFolderInfo) error
}
