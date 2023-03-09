/*
EmployeeのDTO
*/
package dto

import (
	"time"

	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/employee"
	"gorm.io/gorm"
)

type Employee struct {
	ID         uint           `gorm:"primary_key"`
	Name       string         `gorm:"not null"`
	HourlyWage uint           `gorm:"not null"`
	CreateAt   time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time      `gorm:"default:current_timestamp"`
	DeletedAt  gorm.DeletedAt `gorm:"default:null"`
}

type Employees []Employee

func (employee Employee) ConvertToModel() *entity.Employee {
	return &entity.Employee{
		ID:         entity.ID(employee.ID),
		Name:       entity.Name(employee.Name),
		HourlyWage: entity.HourlyWage(employee.HourlyWage),
	}
}

func (employees Employees) ConvertToModel() *entity.Employees {
	result := make(entity.Employees, len(employees))

	for idx, employee := range employees {
		employeeEntity := entity.Employee{
			ID:         entity.ID(employee.ID),
			Name:       entity.Name(employee.Name),
			HourlyWage: entity.HourlyWage(employee.HourlyWage),
		}
		result[idx] = employeeEntity
	}

	return &result
}
