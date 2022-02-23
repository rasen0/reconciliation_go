package handle

import (
	"financial/m/v1/src/log"
	"fmt"
)

type reconciliationTool struct {
	excelData      *ExcelData
	reconciliation *Reconciliation
}

func HandleReconciliation(path string) {
	recon := reconciliationTool{
		excelData:      &ExcelData{},
		reconciliation: NewReconciliation(),
	}
	recon.HandleData(path)
}

func (r *reconciliationTool) HandleData(path string) error {
	r.excelData.Init(path)
	r.excelData.ReadSheet()
	// if err != nil {
	// 	return err
	// }
	r.reconciliation.RowNum = r.excelData.RowNum
	r.reconciliation.ColNum = r.excelData.ColNum
	r.reconciliation.SheetName = r.excelData.Sheets[0]
	r.reconciliation.EFile = r.excelData.EFile
	r.reconciliation.Handle()
	r.Color()

	r.excelData.EFile.SaveAs(r.excelData.FilePath)
	return nil
}

func (r *reconciliationTool) Color() {
	for _, companyStat := range r.reconciliation.CompanyList {
		for _, tmpIdx := range companyStat.RowIdxs {
			if companyStat.State != 1 {
				continue
			}
			rowIdx := fmt.Sprintf("%d", tmpIdx+2)
			style, err := r.excelData.EFile.NewStyle(`{"fill":{"type":"gradient","color":["#7FFFAA","#7FFFAA"],"shading":1}}`)
			if err != nil {
				log.Errorf("colos:%s\n", err)
			}
			log.Debugf("row: %s\n", locationCol["5"]+rowIdx)
			r.excelData.EFile.SetCellStyle(r.excelData.Sheets[0], locationCol["5"]+rowIdx, locationCol["5"]+rowIdx, style)
		}
	}

}
