package ScreenPrinters

import (
	"FileSystemTree4.mod/FileSystemTree4/infrastructure/config"
	"fmt"
)

type ScreenPrinterColor struct {
	FilesColor    string
	FoldersColor  string
	FilesizeColor string
	ErrorColor    string
	ColorReset 		string


}

func (sp ScreenPrinterColor) Print(s string) {
	fmt.Print(s)
}

func (sp ScreenPrinterColor) PrintFolderName(s string) {
	fmt.Print(sp.FoldersColor)
	fmt.Print("["+s+"]")
	fmt.Print(sp.ColorReset)

}

func (sp ScreenPrinterColor) PrintFileName(s string) {
	fmt.Print(sp.FilesColor)
	fmt.Print(s)
	fmt.Print(sp.ColorReset)
}

func (sp ScreenPrinterColor) PrintFileSize(s string) {
	fmt.Print(sp.FilesizeColor)
	fmt.Print(s)
	fmt.Print(sp.ColorReset)
}

func (sp ScreenPrinterColor) PrintError(s string) {
	fmt.Print(sp.ErrorColor)
	fmt.Print(s)
	fmt.Print(sp.ColorReset)
}

func NewScreenPrinterColor(c *config.Config) *ScreenPrinterColor {
	colorPrinter:=ScreenPrinterColor{}

	colorPrinter.FilesizeColor = GetColorByName(c.FilesColor)
	colorPrinter.FoldersColor = GetColorByName(c.FoldersColor)
	colorPrinter.FilesizeColor = GetColorByName(c.FileStatsColor)
	colorPrinter.ErrorColor = GetColorByName(c.WarningColor)
	colorPrinter.ColorReset= "\033[0m"

	return &colorPrinter
}

func GetColorByName(s string) string {

	var (
		colorReset = "\033[0m"

		colorRed    = "\033[31m"
		colorGreen  = "\033[32m"
		colorYellow = "\033[33m"
		colorBlue   = "\033[34m"
		colorPurple = "\033[35m"
		colorCyan   = "\033[36m"
		colorWhite  = "\033[37m")

		switch s {
	case "red":
		return colorRed
	case "green":
		return colorGreen
	case "blue":
		return colorBlue
	case "cyan":
		return colorCyan
	case "purple":
		return colorPurple
	case "yellow":
		return colorYellow
	case "white":
		return colorWhite
	default:
		return colorReset

	}
}