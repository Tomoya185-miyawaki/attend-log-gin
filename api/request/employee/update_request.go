/*
フロントからのリクエストのstruct用のパッケージ
*/
package employee

type EmployeeUpdateRequest struct {
	Name       string `json:"name" binding:"required,max=255"`
	HourlyWage string `json:"hourlyWage" binding:"required"`
}
