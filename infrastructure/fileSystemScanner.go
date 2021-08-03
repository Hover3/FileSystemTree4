package infrastructure

import (
	"FileSystemTree4.mod/FileSystemTree4/domain/folderTree"
	"os"
)

type FileSystemScanner struct {

}

func ExtractFolderContent (f *folderTree.FolderInfo) (
	treeFolderInfo []*folderTree.FolderInfo,
		fileInfo []*folderTree.FileInfo,
		err error){

	FileSystemObjects, err := os.ReadDir(f.GetAbsolutePath())
	if err != nil {
		f.CantAccess = true
		return
	}
	for i := range FileSystemObjects {
		fso:=FileSystemObjects[i]
		if fso.IsDir() {
			//adding directory
			f.AddSubfolderItem(fso.Name())
		} else {
			//adding file
			tempFile := folderTree.FileInfo{FileName: fso.Name()}
			fileInformation, err := fso.Info()
			if err != nil {
				tempFile.CantAccess = true
			} else {
				tempFile.FileSize = fileInformation.Size()
			}
			f.Files = append(f.Files, &tempFile)
		}
	}
	f.IsScanned = true
	err=nil
	return
}