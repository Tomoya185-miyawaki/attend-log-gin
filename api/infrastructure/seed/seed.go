/*
シーダー用パッケージ（CLIからのみ実行）
*/
package main

import (
	"fmt"
	"os"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var Admin = &dto.Admin{
	ID:       1,
	Name:     "管理者",
	Email:    "admin@example.com",
	Password: "$2a$12$jA9fyFy1qfh379BEfaJM2uGPe9EnqrEZnREv1iaiX8nyCmQz69ERK", // passwordをハッシュ化したもの
}

var Employee = &dto.Employee{
	ID:         1,
	Name:       "従業員",
	HourlyWage: 1000,
}

var Employee2 = &dto.Employee{
	ID:         2,
	Name:       "従業員2",
	HourlyWage: 1000,
}

var Employee3 = &dto.Employee{
	ID:         3,
	Name:       "従業員3",
	HourlyWage: 1000,
}

// シーダーを実行
func main() {
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

	if env != "production" {
		// マイグレーション
		db.
			DropTable(&dto.Admin{}).
			DropTable(&dto.Employee{})
		db.
			AutoMigrate(&dto.Admin{}).
			AutoMigrate(&dto.Employee{})
		// シーダーを実行
		db.
			Create(Admin).
			Create(Employee).
			Create(Employee2).
			Create(Employee3)
	}
}
