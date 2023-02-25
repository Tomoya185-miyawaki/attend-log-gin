/*
Employee用のインターフェース
*/
package IDENT

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
)

type EmployeeInterface interface {
	// 従業員を取得する（ページネーションあり）
	FetchEmployees(page int, limit int, offset int) (*dto.Employees, error, int64)
}
