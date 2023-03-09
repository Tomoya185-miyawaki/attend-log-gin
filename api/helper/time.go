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

func SetLocation(timezone string) {
	zone, err := time.LoadLocation(timezone)
	if err != nil {
		log.Error("タイムゾーンの設定に失敗しました")
	}
	time.Local = zone
}

func TimeToString(t time.Time) string {
	return t.Format(layout)
}

func TimeToStringPointer(t *time.Time) string {
	return t.Format(layout)
}

func TimeToStringDuration(t time.Time) string {
	return t.Format(layoutDuration)
}

func TimeToStringDurationPointer(t *time.Time) string {
	return t.Format(layoutDuration)
}

func StringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}
