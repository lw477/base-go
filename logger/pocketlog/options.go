package pocketlog

import "io"

// Option 定义Logger的函数式选项类型
type Option func(logger *Logger)

// WithOutput 返回一个Option，用来设置日志输出目标
func WithOutput(w io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = w
	}
}
