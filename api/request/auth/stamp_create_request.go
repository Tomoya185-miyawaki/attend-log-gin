/*
フロントからのリクエストのstruct用のパッケージ
*/
package auth

import "time"

type StampCreateRequest struct {
	EmployeeId     int       `json:"employeeId" binding:"required"`
	Status         int       `json:"status" binding:"required"`
	StampStartDate time.Time `json:"stamp_start_date" binding:"required"`
	StampEndDate   time.Time `json:"stamp_end_date" binding:"required"`
}
