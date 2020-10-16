package internal

import (
	"fmt"

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

func ReadExcelBook(filename string) ([]Cell, error) {
	cells := []Cell{}

	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Cloud not open %v: %v", filename, err)
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

	return cells, nil
}
