package folderTree



type FolderInfo struct {
	Parent     *FolderInfo
	SubFolders []*FolderInfo
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

func NewRootItem(absolutePath string) *FolderInfo {
	return &FolderInfo{
		Parent:     nil,
		SubFolders: make([]*FolderInfo, 0),
		Files:      make([]*FileInfo, 0),
		IsScanned:  false,
		FolderName: absolutePath,
		CantAccess: false,
	}
}

func (f *FolderInfo) AddSubfolderItem(folderName string) *FolderInfo {

	tmp := &FolderInfo{
		Parent:     f,
		SubFolders: make([]*FolderInfo, 0),
		Files:      make([]*FileInfo, 0),
		IsScanned:  false,
		CantAccess: false,
		FolderName: folderName,
	}
	f.SubFolders=append(f.SubFolders, tmp)
	return tmp
}

func (f *FolderInfo) GetAbsolutePath() string {
	if f.Parent == nil {
		return f.FolderName
	} else {
		return f.Parent.GetAbsolutePath() + "\\" + f.FolderName
	}
}

func (f *FolderInfo) ScanRecurrent(fileSystemScanner FolderContentExtractor) {
	f.ScanFolderContent(fileSystemScanner)
	if f.CantAccess != true {
		for i := range f.SubFolders {
			f.SubFolders[i].ScanRecurrent(fileSystemScanner)
		}
	}
}

func (f *FolderInfo) ScanFolderContent(fileSystemScanner FolderContentExtractor) {
	var err error
	err = fileSystemScanner.ExtractFolderContent(f)
	if err != nil {
		f.CantAccess = true
	}

}
