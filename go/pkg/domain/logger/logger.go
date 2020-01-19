package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	debug = "debug"
	info  = "info"
	warn  = "warn"
	err   = "error"
	fatal = "fatal"
)

// Debug debug用のログ
func Debug(v ...interface{}) {
	os.Stdout.Write([]byte(logText(debug, v...)))
}

// Info infoレベルのログを出力。主にサーバの情報
func Info(v ...interface{}) {
	os.Stdout.Write([]byte(logText(info, v...)))
}

// Warn warnレベルのログを出力。
func Warn(v ...interface{}) {
	os.Stdout.Write([]byte(logText(warn, v...)))
}

// Error エラーの出力。サーバのエラーとか
func Error(v ...interface{}) {
	os.Stdout.Write([]byte(logText(err, v...)))
}

// Fatal エラー出力。アプリケーションを終了させます
func Fatal(v ...interface{}) {
	os.Stdout.Write([]byte(logText(fatal, v...)))
	os.Exit(1)
}

func logText(logLevel string, v ...interface{}) string {
	return fmt.Sprintf("%s [%s] %s\n", time.Now().Format("2006/01/02 15:04:05"), logLevel, fmt.Sprint(v...))
}
