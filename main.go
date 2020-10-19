package main

import (
	"fmt"
	"os"
	"sync"

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

	var wg sync.WaitGroup
	wg.Add(2)

	origin := make(chan []internal.Cell, 1)
	target := make(chan []internal.Cell, 1)
	go internal.ReadExcelBook(&wg, origin, *originFile)
	go internal.ReadExcelBook(&wg, target, *targeFile)
	wg.Wait()

	internal.Diff(origin, target)
}
