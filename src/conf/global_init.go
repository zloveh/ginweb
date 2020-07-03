package conf

import (
	"ginweb/src/util"
)

// 初始化日志
func InitLogger(cl LogConfig) {
	util.InitLooger(
		cl.Level,
		cl.Dir,
		cl.Filename,
		cl.ReserveNum,
		cl.Suffix,
		cl.Console,
		cl.Colorfull)
}
