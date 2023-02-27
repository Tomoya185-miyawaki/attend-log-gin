/*
employeesテーブルを扱うための実態
*/
package repository

import (
	"errors"

	db "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	IDENT "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/interface"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/employee"
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
	if err := db.Limit(limit).Offset(offset).Find(employees).Error; err != nil {
		return nil, errors.New("従業員を取得できませんでした"), total
	}
	return employees, nil, total
}

func (er *employeeRepository) FetchEmployeeById(
	id string,
) (*dto.Employee, error) {
	db := db.GetDB()
	employee := &dto.Employee{}

	if err := db.First(&employee, id).Error; err != nil {
		return nil, errors.New("従業員を取得できませんでした")
	}
	return employee, nil
}

func (er *employeeRepository) Create(request *employee.EmployeeCreateRequest, hourlyWage uint) error {
	db := db.GetDB()
	employee := &dto.Employee{Name: request.Name, HourlyWage: hourlyWage}

	if err := db.Create(&employee).Error; err != nil {
		return errors.New("従業員を取得できませんでした")
	}
	return nil
}
