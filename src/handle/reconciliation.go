package handle

import (
	"financial/m/v1/src/entity"
	"financial/m/v1/src/log"
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Reconciliation struct {
	RowNum      int
	ColNum      int
	SheetName   string
	EFile       *excelize.File
	RowDataList []entity.ReconciliationData // 每行数据
	CompanyList []entity.ReconciliationStat // 单位核对状态
	ReconMap    map[string]int              // 单位名称与下标
}

func NewReconciliation() *Reconciliation {
	return &Reconciliation{
		RowDataList: make([]entity.ReconciliationData, 0),
		CompanyList: make([]entity.ReconciliationStat, 0),
		ReconMap:    make(map[string]int),
	}
}

func (r *Reconciliation) ReadData() {
	log.Debugf("colnum:%d.  rownum:%d", r.ColNum, r.RowNum)
	// deduction := 1
	for ri := 1; ri <= r.RowNum; ri++ {
		cell1, _ := r.EFile.GetCellValue(r.SheetName, locationCol["1"]+fmt.Sprint(ri))
		cell2, _ := r.EFile.GetCellValue(r.SheetName, locationCol["2"]+fmt.Sprint(ri))
		cell3, _ := r.EFile.GetCellValue(r.SheetName, locationCol["3"]+fmt.Sprint(ri))
		cell4, _ := r.EFile.GetCellValue(r.SheetName, locationCol["4"]+fmt.Sprint(ri))
		cell5, _ := r.EFile.GetCellValue(r.SheetName, locationCol["5"]+fmt.Sprint(ri))
		cell6, _ := r.EFile.GetCellValue(r.SheetName, locationCol["6"]+fmt.Sprint(ri))
		cell7, _ := r.EFile.GetCellValue(r.SheetName, locationCol["7"]+fmt.Sprint(ri))
		// if cell5 == "摘要" || cell5 == "" {
		// 	deduction++
		// 	continue
		// }
		// voucherNum, err0 := strconv.Atoi(cell4)
		debit, err1 := strconv.ParseFloat(cell6, 32)
		credit, err2 := strconv.ParseFloat(cell7, 32)

		log.Debugf("row： |%s|  |%s|  |%s| |%s|  |%s|  |%s|  |%s|.debit:%f. credit:%f\n", cell1, cell2, cell3, cell4, cell5, cell6, cell7, debit, credit)
		r.RowDataList = append(r.RowDataList, entity.ReconciliationData{
			Year:       cell1,
			Month:      cell2,
			Day:        cell3,
			VoucherNum: cell4,
			Abstract:   cell5,
			Debit:      debit,
			Credit:     credit,
		})
		if err1 != nil && err2 != nil {
			// deduction++
			continue
		}
		if debit == 0 && credit == 0 && cell5 == "" {
			// deduction++
			continue
		}

		companyName := obtainName(cell5)
		if companyName == "" {
			log.Debugf("name is less %v\n", cell5)
			continue
		}
		// companyName, err := util.RegexpParentheses(cell5)
		// if err != nil {
		// 	log.Errorf("RegexpParentheses error %v\n", err)
		// }
		log.Debugf("name:%s\n", companyName)
		listIdx := ri
		if index, ok := r.ReconMap[companyName]; !ok {
			index = len(r.CompanyList)
			r.ReconMap[companyName] = index
			log.Debugf("name:%s. index:%d. rownum = %d\n", companyName, index, listIdx)
			r.CompanyList = append(r.CompanyList, entity.ReconciliationStat{
				Name:    companyName,
				RowIdxs: []int{listIdx},
			})
		} else {
			reconciliationStat := r.CompanyList[index]
			log.Debugf("name:%s. index:%d. rownum = %d\n", companyName, index, listIdx)
			reconciliationStat.RowIdxs = append(reconciliationStat.RowIdxs, listIdx)
			r.CompanyList[index] = reconciliationStat
		}
	}
	log.Debugf("RowDataList len:%d.", len(r.RowDataList))
	log.Debug("=============================")
}

func (r *Reconciliation) Match() {
	for n, companyVal := range r.CompanyList {
		for _, rowIdx := range companyVal.RowIdxs {
			// if strings.Contains(companyVal.Name, "中国农业银行股份有限公司郴州分行") {
			// 	fmt.Println()
			// }
			// 对比借贷
			log.Debugf("name:%s. rowIdx:%d", companyVal.Name, rowIdx)
			log.Debugf("name:%s.  remain:%f.  debit:%f.   Credit:%f", companyVal.Name, companyVal.Remain, r.RowDataList[rowIdx-1].Debit, r.RowDataList[rowIdx-1].Credit)
			companyVal.Remain += r.RowDataList[rowIdx-1].Debit - r.RowDataList[rowIdx-1].Credit
		}
		if companyVal.Remain == 0 {
			r.CompanyList[n].State = entity.ZeroState
			for _, i := range companyVal.RowIdxs {
				log.Debugf("list  %d\n", i)
			}
			log.Debugf("company:%s, state:%d, \n", companyVal.Name, companyVal.State)
		}
	}
	log.Debug("========================")
}

// func (r *Reconciliation) showResult() {
// 	for i, companyVal := range r.CompanyList {
// 		for _, rowIdx := range companyVal.RowIdxs {
// 			// 对比借贷
// 			if companyVal.State
// 		}
// 	}
// }

func (r *Reconciliation) Handle() {
	r.ReadData()
	r.Match()
}

// obtain company name
func obtainName(str string) string {
	if len(str) < 3 {
		return ""
	}
	var leftIdx int
	idxFirst := strings.Index(str, "（")
	if idxFirst >= 0 {
		leftIdx = idxFirst + 3
	}
	idxFirst2 := strings.Index(str, "(")
	if idxFirst < 0 || (leftIdx >= 0 && idxFirst2 >= 0 && idxFirst2 < idxFirst) {
		leftIdx = idxFirst2 + 1
	}
	if leftIdx < 0 {
		return ""
	}
	var right int
	idxSecond := strings.LastIndex(str, "）")
	if idxSecond > 0 {
		right = idxSecond
	}
	idxSecond2 := strings.LastIndex(str, ")")
	if idxSecond2 > 0 && idxSecond2 > right {
		right = idxSecond2
	}
	if right < 0 {
		return ""
	}
	return str[leftIdx:right]
}

// func (r *Reconciliation) ReadData(rawData [][]string) {
// 	for i, row := range rawData {
// 		voucherNum, err0 := strconv.Atoi(row[3])
// 		debit, err1 := strconv.ParseInt(row[5], 10, 64)
// 		credit, err2 := strconv.ParseInt(row[6], 10, 64)

// 		fmt.Printf("row： |%s|  |%s|  |%s| |%s|  |%s|  |%s|  |%s|\n", row[0], row[1], row[2], row[3], row[4], row[5], row[6])
// 		if err0 != nil && err1 != nil && err2 != nil {
// 			continue
// 		}

// 		companyName := obtainName(row[4])
// 		fmt.Println("name:", companyName)
// 		if index, ok := r.ReconMap[companyName]; !ok {
// 			index = len(r.CompanyList)
// 			r.ReconMap[companyName] = index
// 			r.CompanyList = append(r.CompanyList, entity.ReconciliationStat{
// 				Name:    companyName,
// 				RowIdxs: []int{i},
// 			})
// 		} else {
// 			reconciliationStat := r.CompanyList[index]
// 			reconciliationStat.RowIdxs = append(reconciliationStat.RowIdxs, i)
// 		}
// 		r.RowDataList = append(r.RowDataList, entity.ReconciliationData{
// 			Year:       row[0],
// 			Month:      row[1],
// 			Day:        row[2],
// 			VoucherNum: voucherNum,
// 			Abstract:   row[4],
// 			Debit:      debit,
// 			Credit:     credit,
// 		})
// 	}

// }
