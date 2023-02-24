/*
DB用パッケージ
*/
package infrastructure

import (
	"fmt"
	"os"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/seed"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// DBの初期化
func Init() *gorm.DB {
	env := os.Getenv("ENV")
	if env != "production" {
		env = "development"
	}
	godotenv.Load(".env." + env)
	godotenv.Load()

	db, err = gorm.Open("mysql", os.Getenv("DB_CONNECT"))
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	// マイグレーション
	autoMigrate()
	if env != "production" {
		// シーダーを実行
		execSeeds()
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

// マイグレーション
func autoMigrate() {
	db.
		AutoMigrate(&dto.Admin{})
}

func execSeeds() {
	db.
		Create(seed.Admin)
}
