package pocketlog

// Level 代表日志的重要级别
type Level byte

const (
	// LevelDebug 调试级别，最低等级，用于开发调试
	LevelDebug Level = iota

	// LevelInfo 普通业务信息
	LevelInfo

	// LevelError 错误级别，最高优先级
	LevelError

	// 数值大小决定阈值逻辑：数字越大级别越严重。
)

