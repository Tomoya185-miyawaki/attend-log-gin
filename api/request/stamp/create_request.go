/*
フロントからのリクエストのstruct用のパッケージ
*/
package stamp

type StampCreateRequest struct {
	EmployeeId int    `json:"employeeId" binding:"required"`
	Status     int    `json:"status" binding:"required"`
	Today      string `json:"today" binding:"required"`
}
