package ScreenPrinters

import "fmt"

type ScreenPrinterMonochrome struct {

}

func (sp ScreenPrinterMonochrome) Print (s string) {
	fmt.Print(s)
}

func (sp ScreenPrinterMonochrome) PrintFolderName (s string) {
	fmt.Print("["+s+"]")
}

func (sp ScreenPrinterMonochrome) PrintFileName (s string) {
	fmt.Print(s)
}

func (sp ScreenPrinterMonochrome) PrintFileSize (s string) {
	fmt.Print(s)
}

func (sp ScreenPrinterMonochrome) PrintError (s string) {
	fmt.Print(s)
}




