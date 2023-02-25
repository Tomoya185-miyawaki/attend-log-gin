/*
EmployeeのDTO
*/
package dto

import (
	"time"
)

type Employee struct {
	ID         uint      `gorm:"primary_key"`
	Name       string    `gorm:"not null"`
	HourlyWage uint      `gorm:"not null"`
	CreateAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp"`
}
