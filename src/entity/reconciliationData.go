package entity

type ReconciliationData struct {
	Year       string
	Month      string
	Day        string
	VoucherNum int //凭证号
	Abstract   string
	Debit      int64 //借方
	Credit     int64 //贷方
}
