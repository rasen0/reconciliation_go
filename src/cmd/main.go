package main

import (
	"financial/m/v1/src/configure"
	"financial/m/v1/src/handle"
	"financial/m/v1/src/log"
)

func main() {
	// todo 读取配置文件
	cfg := configure.New()
	// todo 加载日志系统
	log.Init(cfg.LogPath)
	log.Debug("...start...")
	// 读取文件,导出结果
	handle.HandleReconciliation("核对账.xlsx")
}
