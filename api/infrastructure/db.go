/*
DB用パッケージ
*/
package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

// DBの初期化
func Init() *gorm.DB {
	env := os.Getenv("ENV")
	if env != "production" {
		env = "development"
	}
	godotenv.Load(".env." + env)
	godotenv.Load()

	db, err := gorm.Open("mysql", os.Getenv("DB_CONNECT"))
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
