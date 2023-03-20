/*
DB用パッケージ
*/
package infrastructure

import (
	"os"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

// DBの初期化
func Init() *gorm.DB {
	// env読み込み
	helper.GetEnv()

	db, err = gorm.Open("mysql", os.Getenv("DB_CONNECT"))
	if err != nil {
		log.Warn("db init error: ", err)
	}

	// マイグレーション
	autoMigrate()

	return db
}

func GetDB() *gorm.DB {
	return db
}

// マイグレーション
func autoMigrate() {
	db.
		AutoMigrate(&dto.Admin{}).
		AutoMigrate(&dto.Employee{})
}
