/*
Employee用のインターフェース
*/
package IDENT

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/employee"
)

type EmployeeInterface interface {
	// 従業員を取得する（ページネーションなし）
	FetchEmployees() (*dto.Employees, error)
	// 従業員を取得する（ページネーションあり）
	FetchEmployeesByPaginate(page int, limit int, offset int) (*dto.Employees, error, int64)
	// IDをキーにして、従業員を取得する
	FetchEmployeeById(id string) (*dto.Employee, error)
	// IDをキーにして、複数従業員を取得する
	FetchEmployeeByIds(ids []uint) (*dto.Employees, error)
	// 従業員を作成する
	Create(request *employee.EmployeeCreateRequest, hourlyWage uint) error
	// 従業員を更新する
	Update(id string, name string, hourlyWage uint) error
	// 従業員を削除する
	Delete(id string) error
}
