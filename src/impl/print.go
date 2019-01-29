package impl

import (
	"fmt"
	"time"
	"errors"
)

//日志模式
const (
	FlushLogAtMoment  = 1
	FlushLogInterval  = 2
	FLushLogSpecified = 3
)

//日志级别
const (
	Debug   = 1
	Info    = 2
	Notice  = 3
	Warning = 4
	Error   = 5
	RunTine = 6
)

type Logger struct {
}

var SystemOutput *Logger

/**
 *	系统输出
 */
func init() {
	SystemOutput = &Logger{}
}

func (l *Logger) Ding() {

}

/**
 *	格式化错误信息
 *	用于需要将自定义错误生成error类型时使用 会自动调用打印功能
 */
func (l *Logger) FormatLog(s ...interface{}) error {
	str := ""
	for _, v := range s {
		str += fmt.Sprint(v) + " "
	}
	go l.Debug(str)
	return errors.New(str)
}

/**
 *	格式化错误信息 通用方法
 */
func FormatLog(s ...interface{}) error {
	str := ""
	for _, v := range s {
		str += fmt.Sprint(v) + " "
	}
	logger := &Logger{}
	logger.Error(str)
	return errors.New(str)
}

/**
 *	Error类型日志,会自动记录出发error的文件行号
 */
func (l *Logger) Error(s ...interface{}) {
	l.format(Error, s...)
}

/**
 *	Warning类型日志 粉色
 */
func (l *Logger) Warning(s ...interface{}) {
	l.format(Warning, s...)
}

/**
 *	Notice类型日志 黄色
 */
func (l *Logger) Notice(s ...interface{}) {
	l.format(Notice, s...)
}

/**
 *	Info类型日志 蓝色
 */
func (l *Logger) Info(s ...interface{}) {
	l.format(Info, s...)
}

/**
 *	Debug类型日志 绿色
 */
func (l *Logger) Debug(s ...interface{}) {
	l.format(Debug, s...)
}

/**
 *	goruntine专用
 */
func (l *Logger) Run(s ...interface{}) {
	l.format(RunTine, s...)
}

/**
 *	输出日志文件信息
 */
func (l *Logger) format(level int, s ...interface{}) {
	var loginfo, tags string
	switch level {
	case Debug:
		tags = "[1;1;32m[Debug]"
	case Info:
		tags = "[1;1;34m[Info]"
	case Notice:
		tags = "[1;1;33m[Notice]"
	case Warning:
		tags = "[1;1;35m[Warning]"
	case Error:
		tags = "[1;1;31m[Error]"
	default:
		tags = "[1;1;36m[Run]"
	}
	for _, v := range s {
		loginfo = loginfo + " " + fmt.Sprint(v)
	}
	timeStr := time.Now().Format("06-01-02 15:04:05")
	fmt.Printf("%c%s%c[0m %s%s\n", 0x1B, tags, 0x1B, timeStr, loginfo)
}
