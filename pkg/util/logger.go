package util

import (
	"fmt"
	"os"
	"time"
)

const (
	// LevelError 错误
	LevelError = iota
	// LevelWarning 警告
	LevelWarning
	// LevelInformational 提示
	LevelInformational
	// LevelDebug 除错
	LevelDebug
)

var logger *Logger

type Logger struct {
	level int
}

// Println 打印
func (ll *Logger) Println(msg string) {
	fmt.Printf("s s", time.Now().Format("2006-01-01 15:01:05 -0700"), msg)
}

// Panic 极端错误
func (ll *Logger) Panic(format string, v ...interface{}) {
	// 初始判断，不在范围之内则直接返回
	if ll.level < LevelError {
		return
	}
	// msg进行拼接
	msg := fmt.Sprintf("[Panic]"+format, v...)
	ll.Println(msg)
	os.Exit(0)
}

// Error 错误
func (ll *Logger) Error(format string, v ...interface{}) {
	if ll.level < LevelError {
		return
	}
	msg := fmt.Sprintf("[Error]"+format, v...)
	ll.Println(msg)
}

// Warning 警告
func (ll *Logger) Warning(format string, v ...interface{}) {
	if ll.level < LevelWarning {
		return
	}
	msg := fmt.Sprintf("[Waarning]"+format, v...)
	ll.Println(msg)
}

// Info 信息
func (ll *Logger) Info(format string, v ...interface{}) {
	if ll.level < LevelInformational {
		return
	}
	msg := fmt.Sprintf("[InforMation]"+format, v...)
	ll.Println(msg)
}

// Debug 校验
func (ll *Logger) Debug(format string, v ...interface{}) {
	if ll.level < LevelDebug {
		return
	}
	msg := fmt.Sprintf("[Debug]"+format, v...)
	ll.Println(msg)
}

// BuildLogger 构建logger
func BuildLogger(level string) {
	// 初始为LevelError
	// 根据传递进来的level进行初始日志对象
	initLevel := LevelError
	switch level {
	case "error":
		initLevel = LevelError
	case "warning":
		initLevel = LevelWarning
	case "info":
		initLevel = LevelInformational
	case "debug":
		initLevel = LevelDebug
	}
	l := Logger{
		level: initLevel,
	}
	logger = &l
}

// log 返回日志对象
func Log() *Logger {
	// 如果logger为空，则新构建一个并返回，初始level为LevelDebug
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}
	return logger
}
