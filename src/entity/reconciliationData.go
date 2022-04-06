package entity

type ReconciliationData struct {
	Year       string
	Month      string
	Day        string
	VoucherNum string //凭证号
	Abstract   string
	Debit      float64 //借方
	Credit     float64 //贷方
}
