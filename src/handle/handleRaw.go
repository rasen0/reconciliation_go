package handle

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func HandleRaw() {
	f, err := excelize.OpenFile("体检测试0.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	var sheets = f.GetSheetList()
	for i, v := range sheets {
		fmt.Printf("v:%s i:%d\n", v, i)
	}
	fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()

	}
}
