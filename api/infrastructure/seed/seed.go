/*
シーダー用パッケージ（CLIからのみ実行）
*/
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
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

var Stamp1 = &dto.Stamp{
	ID:             1,
	EmployeeID:     1,
	Status:         1,
	StampStartDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 0, 0, 0, time.Local),
}

var Stamp2 = &dto.Stamp{
	ID:             2,
	EmployeeID:     1,
	Status:         2,
	StampStartDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 3, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 4, 0, 0, 0, time.Local),
}

// シーダーを実行
func main() {
	env := os.Getenv("ENV")
	if env != "production" {
		env = "development"
	}
	godotenv.Load(".env." + env)
	godotenv.Load()

	// タイムゾーンの設定
	helper.SetLocation("Asia/Tokyo")

	db, err := gorm.Open("mysql", os.Getenv("DB_CONNECT"))
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	if env != "production" {
		// マイグレーション
		db.
			DropTable(&dto.Admin{}).
			DropTable(&dto.Employee{}).
			DropTable(&dto.Stamp{})
		db.
			AutoMigrate(&dto.Admin{}).
			AutoMigrate(&dto.Employee{}).
			AutoMigrate(&dto.Stamp{})
		// シーダーを実行
		db.
			Create(Admin).
			Create(Employee).
			Create(Employee2).
			Create(Employee3).
			Create(Stamp1).
			Create(Stamp2)
	}
}
