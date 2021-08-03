package interfaces

import "FileSystemTree4.mod/FileSystemTree4/domain/folderTree"

type FolderContentExtractor interface {
	ExtractFolderContent (f *folderTree.FolderInfo)  error
}
