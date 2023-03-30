/*
フロントからのリクエストのstruct用のパッケージ
*/
package stamp

import "time"

type StampUpdateRequest struct {
	Id             int       `json:"id" binding:"required"`
	Status         int       `json:"status" binding:"required"`
	StampStartDate time.Time `json:"stamp_start_date" binding:"required"`
	StampEndDate   time.Time `json:"stamp_end_date" binding:"required"`
}
