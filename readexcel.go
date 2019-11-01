package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
)

func main() {
    excelFileName := "c:\\temp\\Бренды различные названия.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        
    }
    
    for _, sheet := range xlFile.Sheets {
        sheetName := sheet.Name
        fmt.Printf(sheetName)
    }

    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
                text := cell.String()
                fmt.Printf("%s\n", text)
            }
        }
    }

}