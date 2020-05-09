package logger

import (
	"fmt"
	"homepage/conf"
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

var logfile *os.File

func SetUpLogfile(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	logfile = file
	return nil
}

func CloseLogfile() error {
	return logfile.Close()
}

// Debug debug用のログ
func Debug(v ...interface{}) {
	os.Stdout.Write([]byte(logText(debug, v...)))
}

// Info infoレベルのログを出力。主にサーバの情報
func Info(v ...interface{}) {
	os.Stdout.Write([]byte(logText(info, v...)))
	logfile.Write([]byte(logText(info, v...)))
}

// Warn warnレベルのログを出力。
func Warn(v ...interface{}) {
	os.Stdout.Write([]byte(logText(warn, v...)))
	logfile.Write([]byte(logText(warn, v...)))
}

// Error エラーの出力。サーバのエラーとか
func Error(v ...interface{}) {
	os.Stdout.Write([]byte(logText(err, v...)))
	logfile.Write([]byte(logText(err, v...)))
}

// Fatal エラー出力。アプリケーションを終了させます
func Fatal(v ...interface{}) {
	os.Stdout.Write([]byte(logText(fatal, v...)))
	os.Exit(1)
}

func logText(logLevel string, v ...interface{}) string {
	return fmt.Sprintf("%s [%s] %s\n", time.Now().Format(conf.LogDatetimeFormat), logLevel, fmt.Sprint(v...))
}
