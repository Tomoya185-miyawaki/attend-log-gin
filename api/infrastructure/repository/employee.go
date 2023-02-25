/*
employeesテーブルを扱うための実態
*/
package repository

import (
	"errors"
	"fmt"

	db "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	IDENT "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/interface"
)

type employeeRepository struct{}

func NewEmployeeRepository() IDENT.EmployeeInterface {
	return &employeeRepository{}
}

func (er *employeeRepository) FetchEmployees(
	page int,
	limit int,
	offset int,
) (*dto.Employees, error, int64) {
	db := db.GetDB()
	employees := &dto.Employees{}

	total := db.Find(&employees).RowsAffected
	fmt.Println(limit, offset)
	if err := db.Limit(limit).Offset(offset).Find(employees).Error; err != nil {
		return nil, errors.New("従業員を取得できませんでした"), total
	}
	return employees, nil, total
}
