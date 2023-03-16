/*
time用ミドルウェア
*/
package helper

import (
	"time"

	log "github.com/sirupsen/logrus"
)

var layout = "2006-01-02 15:04:05"
var layoutDuration = "15:04"

func SetAsiaLocation() {
	time.Local = time.FixedZone("Local", 9*60*60) // 日本時間にする
	_, err := time.LoadLocation("Local")
	if err != nil {
		log.Error("タイムゾーンの設定に失敗しました")
	}
}

func TimeToString(t time.Time) string {
	return t.Format(layout)
}

func TimeToStringDuration(t time.Time) string {
	return t.Format(layoutDuration)
}
