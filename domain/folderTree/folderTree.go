package folderTree

type TreeFolderInfo struct {
	Parent     *TreeFolderInfo
	SubFolders []*TreeFolderInfo
	Files      []*FileInfo
	IsScanned  bool
	CantAccess bool
	FolderName string
}

type FileInfo struct {
	FileName   string
	FileSize   int64
	FileExt    string
	CantAccess bool
}

func NewRootItem(absolutePath string) *TreeFolderInfo {
	return &TreeFolderInfo{
		Parent:     nil,
		SubFolders: make([]*TreeFolderInfo, 0),
		Files:      make([]*FileInfo, 0),
		IsScanned:  false,
		FolderName: absolutePath,
		CantAccess: false,
	}
}

func newSubFolderItem(folderName string, parent *TreeFolderInfo) *TreeFolderInfo {
	return &TreeFolderInfo{
		Parent:     parent,
		SubFolders: make([]*TreeFolderInfo, 0),
		Files:      make([]*FileInfo, 0),
		IsScanned:  false,
		CantAccess: false,
		FolderName: folderName,
	}
}

func (f *TreeFolderInfo) GetAbsolutePath() string {
	if f.Parent == nil {
		return f.FolderName
	} else {
		return f.Parent.FolderName + "\\" + f.FolderName
	}
}
