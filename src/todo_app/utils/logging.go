package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 読み書き、ファイルの作成、ファイルの追記
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// ログの書き込み先、標準出力でlogfileへ
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// ログの出力先
	log.SetOutput(multiLogFile)
}
