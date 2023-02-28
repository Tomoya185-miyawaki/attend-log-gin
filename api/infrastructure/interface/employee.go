/*
Employee用のインターフェース
*/
package IDENT

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/employee"
)

type EmployeeInterface interface {
	// 従業員を取得する（ページネーションあり）
	FetchEmployees(page int, limit int, offset int) (*dto.Employees, error, int64)
	// IDをキーにして、従業員を取得する
	FetchEmployeeById(id string) (*dto.Employee, error)
	// 従業員を作成する
	Create(request *employee.EmployeeCreateRequest, hourlyWage uint) error
	// 従業員を更新する
	Update(id string, name string, hourlyWage uint) error
}
