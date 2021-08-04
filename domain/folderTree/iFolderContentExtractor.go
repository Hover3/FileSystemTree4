package folderTree

type FolderContentExtractor interface {
	ExtractFolderContent (f *FolderInfo)  error
}
