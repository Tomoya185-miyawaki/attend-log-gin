/*
ログ用ヘルパー
*/
package helper

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetUpLog() {
	log.SetReportCaller(true)              // 呼び出し元情報
	log.SetFormatter(&log.JSONFormatter{}) // JSON形式で出力
	log.SetOutput(&lumberjack.Logger{
		Filename:   "log/app.log", // ファイル名
		MaxSize:    500,           // ローテーションするファイルサイズ(megabytes)
		MaxBackups: 3,             // 保持する古いログの最大ファイル数
		MaxAge:     365,           // 古いログを保持する日数
		LocalTime:  true,          // バックアップファイルの時刻フォーマットをサーバローカル時間指定
		Compress:   true,          // ローテーションされたファイルのgzip圧縮
	})
}
