package read_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/xuri/excelize/v2"
)

func Test_regexp0(t *testing.T) {
	// str := "df扥恩负二本(扥而dnef覅鞥欧派)"
	// str := "df扥恩负二本（扥欧派）"
	str := "df扥恩负二本(扥欧派）dd"
	// str := "df扥恩负二本(扥欧派"
	// matched, err := regexp.MatchString(".*[\\(](.*)[\\)]", str)
	// fmt.Println(matched, err)
	reg, err := regexp.Compile(".*[\\(（]{1}(.*)[\\)）]{1}.*")
	if err != nil {
		fmt.Println(err)
	}
	sublist := reg.FindStringSubmatch(str)
	if len(sublist) < 2 {
		fmt.Println(len(sublist))
	}
	for _, v := range sublist {
		fmt.Println("v:", v)
	}
	fmt.Println("v:", sublist[1])
}

func Test_read0(t *testing.T) {
	fmt.Println("p:", os.Args[0])
	wd, _ := os.Getwd()
	fmt.Println("p:", wd)
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
	cell, err := f.GetCellValue("Sheet1", "e1")
	if err != nil {
		fmt.Println(err)
		return
	}
	styleId, errr := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Color: []string{"#06FA86"}, Shading: 5},
	})
	if errr != nil {
		fmt.Printf("errr:%s", errr)
	}
	errr = f.SetCellStyle("Sheet1", "e1", "e1", styleId)
	if errr != nil {
		fmt.Printf("errrrr:%s", errr)
	}
	fmt.Println(cell)

	style, err := f.NewStyle(`{"fill":{"type":"gradient","color":["#7FFFAA","#7FFFAA"],"shading":1}}`)
	if err != nil {
		fmt.Println(err)
	}
	_ = f.SetCellStyle("Sheet1", "H9", "H9", style)

	err0 := f.SetCellValue("Sheet1", "H9", "ne")
	if err0 != nil {
		fmt.Println("err0:", err0)
	}
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Printf("  <%s>  ", colCell)
		}
		fmt.Println()
	}
	// f.Save()
	f.SaveAs("体检测试0.xlsx")
	// f.Close()

}
