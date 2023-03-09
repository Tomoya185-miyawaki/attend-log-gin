/*
フロントからのリクエストのstruct用のパッケージ
*/
package stamp

type StampCreateRequest struct {
	EmployeeId int    `json:"EmployeeId" binding:"required"`
	Status     int    `json:"status" binding:"required"`
	StampDate  string `json:"stampDate" binding:"required"`
}
