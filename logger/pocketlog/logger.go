package pocketlog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jinzhu/now"
)

// Logger 日志实例，阈值以上的消息才会输出
type Logger struct {
	threshold Level     // 未导出(小写)：包外不能直接访问
	output    io.Writer // 输出目标，未导出
}

// New 返回一个日志实例，threshold设置日志过滤阈值，可以传入可选配置option
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}
	// 依次执行所有传入的 functional option
	for _, fn := range opts {
		fn(lgr)
	}
	return lgr
}

// logf 内部私有方法，真正执行输出，未导出
func (l *Logger) logf(foramt string, args ...any) {
	_, _ = fmt.Fprintf(l.output, foramt+"\n", args...)
}

// Debugf 调试日志，格式化输出。只有阈值>=Debug才打印
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	l.logf(format, args...)
}

// Infof 信息日志
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	l.logf(format, args...)
}

// Errorf 错误日志
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	l.logf(format, args...)
}


func (l *Logger)WithTimestamp(format string, args ...any){
	var ts time.Time
	ts = time.Now()
	l.logf(format,args...,ts)
}