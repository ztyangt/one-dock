package console

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", 0)
}

var levelColors = map[string]string{
	"DEBUG":    "\033[97;46m", // 白字+青色背景
	"INFO":     "\033[97;42m", // 白字+绿色背景
	"WARNING":  "\033[97;43m", // 白字+黄色背景
	"ERROR":    "\033[97;41m", // 白字+红色背景
	"CRITICAL": "\033[97;45m", // 白字+紫色背景
	"SUCCESS":  "\033[97;42m", // 白字+绿色背景
}

// 获取不同级别的文字颜色（用于时间部分）
var levelTextColors = map[string]string{
	"DEBUG":    "\033[36m", // 青色
	"INFO":     "\033[32m", // 绿色
	"WARNING":  "\033[33m", // 黄色
	"ERROR":    "\033[31m", // 红色
	"CRITICAL": "\033[35m", // 紫色
	"SUCCESS":  "\033[32m", // 绿色
}

var reset = "\033[0m"      // 重置颜色
var whiteText = "\033[97m" // 白色字体

// 自定义日志格式
func format(level string, message string) string {
	// 获取当前时间
	now := time.Now().Format("2006-01-02 15:04:05.000")

	// 获取调用栈信息（可选，如果需要显示文件名和行号）
	//_, file, line, ok := runtime.Caller(2)
	//caller := ""
	//if ok {
	//	// 只显示文件名，不显示完整路径
	//	parts := strings.Split(file, "/")
	//	file = parts[len(parts)-1]
	//	caller = fmt.Sprintf(" [%s:%d]", file, line)
	//}

	// 构建日志格式
	// 格式：[级别] 时间 消息
	// 级别：白字 + 背景色
	// 时间：对应级别的文字颜色
	// 消息：白色字体
	levelPart := fmt.Sprintf("%s[%s]%s", levelColors[level], level, reset)
	timePart := fmt.Sprintf("%s%s%s", levelTextColors[level], now, reset)
	messagePart := fmt.Sprintf("%s%s%s", whiteText, message, reset)

	// 如果有调用信息，添加到消息后面
	//if caller != "" {
	//	messagePart = fmt.Sprintf("%s%s", messagePart, caller)
	//}

	return fmt.Sprintf("%s %s %s", levelPart, timePart, messagePart)
}

// Log 日志输出函数
func Log(message string, level ...string) {
	// 默认级别为INFO
	if len(level) == 0 {
		level = append(level, "INFO")
	}
	logLine := format(level[0], message)
	logger.Println(logLine)
}

// Info 不同级别的日志函数
func Info(message string) {
	Log(message, "INFO")
}

func Infof(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	Log(message, "INFO")
}

func Error(message string) {
	Log(message, "ERROR")
}

func Errorf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	Log(message, "ERROR")
}

func Debug(message string) {
	Log(message, "DEBUG")
}

func Debugf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	Log(message, "DEBUG")
}

func Warn(message string) {
	Log(message, "WARNING")
}

func Warnf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	Log(message, "WARNING")
}

func Critical(message string) {
	Log(message, "CRITICAL")
}

func Criticalf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	Log(message, "CRITICAL")
}

func Success(message string) {
	Log(message, "SUCCESS")
}

func Successf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	Log(message, "SUCCESS")
}
