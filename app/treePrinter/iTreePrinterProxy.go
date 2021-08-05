package treePrinter

type PrintingProxy interface {
	Print (s string)//
	PrintFolderName(s string)
	PrintFileName(s string)
	PrintFileSize(s string)
	PrintError(s string)
}

