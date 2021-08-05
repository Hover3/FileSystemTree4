package treePrinter

import (
	"FileSystemTree4/domain/folderTree"
	"fmt"
)

type TreePrinter struct {
	Proxy  PrintingProxy
	Filter folderTree.ITreeFilter
}

func (tp *TreePrinter) printFiles(f *folderTree.FolderInfo, prefix string) {
	VisibleFilesInThisFolder:=f.CountFilesMatchFilter(tp.Filter)
	PrintedFiles:=0
	for i := range f.Files {


		if tp.Filter.IsFileMatchFilter(f.Files[i]) {
			PrintedFiles++
			s:= "\u251C"
			if PrintedFiles==VisibleFilesInThisFolder {
				s = "\u2514"
			}
			tp.Proxy.Print(prefix + s + "\u2500")
			tp.Proxy.PrintFileName(f.Files[i].FileName)
			tp.Proxy.PrintFileSize(fmt.Sprint("\t", "size:", f.Files[i].FileSize, "\n"))
		}

	}
}

func NewTreePrinter(pp PrintingProxy, filter folderTree.ITreeFilter) *TreePrinter {
	return &TreePrinter{
		Proxy: pp,
		Filter: filter}
}

func (tp *TreePrinter) PrintRecurrent(f *folderTree.FolderInfo, prefix string, last bool) {
	tp.printFolderRow(f,prefix, last)
	for i := 0; i < len(f.SubFolders); i++ {
		s:="  "
		if last==false {	s = "\u2502" + " "}
		isLast:=false
		if i == len(f.SubFolders)-1 && f.CountFilesMatchFilter(tp.Filter)==0 { isLast=true}
		tp.PrintRecurrent(f.SubFolders[i],prefix+s, isLast )
	}
	filesPrefix:="\u2502" + " "
	if last==true {filesPrefix =" " + " "}
	tp.printFiles(f,prefix + filesPrefix)
}

func (tp *TreePrinter) printFolderRow(f *folderTree.FolderInfo, prefix string, last bool) {
	var s string
	if last == true {
		s = "\u2514"
	} else {
		s = "\u251C"
	}
	tp.Proxy.Print(prefix + s + "\u2500")
	tp.Proxy.PrintFolderName(f.FolderName)
	if f.CantAccess == true {
		tp.Proxy.PrintError("\t No access!")
	}
	tp.Proxy.Print("\n")
}



