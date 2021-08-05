package folderTree

type ITreeFilter interface {
	IsFileMatchFilter (f *FileInfo) bool
}
