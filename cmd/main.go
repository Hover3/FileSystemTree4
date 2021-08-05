package main

import (
	"FileSystemTree4/app/treePrinter"
	"FileSystemTree4/domain/folderTree"
	"FileSystemTree4/infrastructure/ScreenPrinters"
	"FileSystemTree4/infrastructure/config"
	"FileSystemTree4/infrastructure/fileSystemScanner"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	conf:=config.New()

	currentDir:=conf.RootDirectory
	if currentDir=="" {
		var err error
		currentDir, err = os.Getwd()
		if err != nil {
			_=currentDir
			panic(err)
		}
	}


	fileSystemScanner:= fileSystemScanner.FileSystemScanner{}

	rootFolder := folderTree.NewRootItem(currentDir)

	rootFolder.ScanRecurrent(&fileSystemScanner)

	//fmt.Println(len(rootFolder.SubFolders), " ", len(rootFolder.Files))
	//fmt.Println(rootFolder.SubFolders[1].FolderName)

	var pp treePrinter.PrintingProxy
	if conf.EnableColorText {
		pp=ScreenPrinters.NewScreenPrinterColor(conf)
	}else {
		pp=ScreenPrinters.ScreenPrinterMonochrome{}
	}


	filter:=treePrinter.FilterExtensions{
		AllowedExtensions: conf.EnabledExtensions,
	}


	tp:=treePrinter.NewTreePrinter(pp, &filter)
	tp.PrintRecurrent(rootFolder,"", true)

}
