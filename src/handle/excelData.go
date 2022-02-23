package handle

import (
	"financial/m/v1/src/log"

	"github.com/xuri/excelize/v2"
)

type ExcelData struct {
	RowNum   int
	ColNum   int
	FilePath string
	Sheets   []string
	EFile    *excelize.File
}

func (e *ExcelData) Init(path string) error {
	var err error
	e.FilePath = path
	e.EFile, err = excelize.OpenFile(path)
	if err != nil {
		if err := e.EFile.Close(); err != nil {
			log.Debug(err.Error())
		}
	}
	return err
}

func (e *ExcelData) GetSheets() []string {
	return e.EFile.GetSheetList()
}

func (e *ExcelData) Close() error {
	return e.EFile.Close()
}

func (e *ExcelData) ReadSheet() {
	e.Sheets = e.EFile.GetSheetList()
	// for i, v := range sheets {
	// 	fmt.Printf("sheets v:%s i:%d\n", v, i)
	// }

	// 获取 Sheet1 上所有单元格
	// rows, err := e.EFile.GetRows(sheets[0])
	r, _ := e.EFile.Rows(e.Sheets[0])
	e.RowNum = r.TotalRows()
	c, _ := e.EFile.Cols(e.Sheets[0])
	e.ColNum = c.TotalCols()
	log.Debugf("r:%d, c:%d \n", e.RowNum, e.ColNum)
	return
}
