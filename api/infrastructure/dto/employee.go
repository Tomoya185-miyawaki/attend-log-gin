/*
EmployeeのDTO
*/
package dto

import (
	"time"

	"github.com/Tomoya185-miyawaki/attend-log-gin/entity"
)

type Employee struct {
	ID         uint      `gorm:"primary_key"`
	Name       string    `gorm:"not null"`
	HourlyWage uint      `gorm:"not null"`
	CreateAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp"`
}

type Employees []Employee

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
