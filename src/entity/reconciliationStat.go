package entity

const (
	InitState = iota
	ZeroState
	DebitState
	CreditState
)

type ReconciliationStat struct {
	State   int    // 剩余值正、负、平
	Remain  int64  // 核对后剩余额度
	Name    string // 单位名
	RowIdxs []int  // 单位行号
}
