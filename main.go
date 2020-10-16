package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/guromityan/go-diffex/internal"
)

const version = "0.9.0"

var (
	cli        = kingpin.New("diffex", "Compare Excel files by cell value.")
	originFile = cli.Arg("origin", "Original file to be compared.").Required().ExistingFile()
	targeFile  = cli.Arg("target", "Target file to be compared.").Required().ExistingFile()
)

func main() {
	cli.Version(version)
	cli.Parse(os.Args[1:])

	if *originFile == "" {
		fmt.Println("Please specify the original file.")
		return
	}

	if *targeFile == "" {
		fmt.Println("Please specify the target file.")
		return
	}

	origin, err := internal.ReadExcelBook(*originFile)
	if err != nil {
		log.Fatalln(err)
	}
	target, _ := internal.ReadExcelBook(*targeFile)
	if err != nil {
		log.Fatalln(err)
	}

	internal.Diff(origin, target)
}
