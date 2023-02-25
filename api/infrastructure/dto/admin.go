/*
AdminユーザーのDTO
*/
package dto

import (
	"time"
)

type Admin struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreateAt  time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
