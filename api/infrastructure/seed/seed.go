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

var jst, _ = time.LoadLocation("Asia/Tokyo")

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

var StampTodayAttend1 = &dto.Stamp{
	ID:             1,
	EmployeeID:     1,
	Status:         1,
	StampStartDate: time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 0, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 9, 0, 0, 0, time.Local),
}

var StampTodayRest1 = &dto.Stamp{
	ID:             2,
	EmployeeID:     1,
	Status:         2,
	StampStartDate: time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 3, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 4, 0, 0, 0, time.Local),
}

var StampYesterdayAttend1 = &dto.Stamp{
	ID:             3,
	EmployeeID:     1,
	Status:         1,
	StampStartDate: time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day()-1, 0, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day()-1, 9, 0, 0, 0, time.Local),
}

var StampYesterdayRest1 = &dto.Stamp{
	ID:             4,
	EmployeeID:     1,
	Status:         2,
	StampStartDate: time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day()-1, 3, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day()-1, 4, 0, 0, 0, time.Local),
}

var StampTodayAttend2 = &dto.Stamp{
	ID:             5,
	EmployeeID:     2,
	Status:         1,
	StampStartDate: time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 0, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 10, 0, 0, 0, time.Local),
}

var StampTodayRest2 = &dto.Stamp{
	ID:             6,
	EmployeeID:     2,
	Status:         2,
	StampStartDate: time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 3, 0, 0, 0, time.Local),
	StampEndDate:   time.Date(time.Now().In(jst).Year(), time.Now().In(jst).Month(), time.Now().In(jst).Day(), 3, 45, 0, 0, time.Local),
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
	helper.SetAsiaLocation()

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
			Create(StampTodayAttend1).
			Create(StampTodayRest1).
			Create(StampYesterdayAttend1).
			Create(StampYesterdayRest1).
			Create(StampTodayAttend2).
			Create(StampTodayRest2)
	}
}
