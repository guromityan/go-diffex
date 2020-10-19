package internal

import (
	"log"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Cell struct {
	Sheet string
	Cell  string
	Value string
}

func NewCell(sheet string, col, row int, val string) Cell {
	cellName, _ := excelize.CoordinatesToCellName(col, row)
	cell := Cell{sheet, cellName, val}
	return cell
}

func ReadExcelBook(wg *sync.WaitGroup, readCells chan []Cell, filename string) {
	cells := []Cell{}

	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	for _, sheet := range f.GetSheetList() {
		rows, _ := f.GetRows(sheet)
		for r, row := range rows {
			for c, val := range row {
				cell := NewCell(sheet, c+1, r+1, val)
				cells = append(cells, cell)
			}
		}
	}

	readCells <- cells
	wg.Done()
}
