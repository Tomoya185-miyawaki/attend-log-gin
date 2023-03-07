/*
レスポンスの型用のパッケージ（従業員）
*/
package employee

import (
	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/employee"
)

type ListResponse struct {
	CurrentPage int              `json:"currentPage"`
	Emplyees    entity.Employees `json:"employees"`
	LastPage    int              `json:"lastPage"`
}

type BadListResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
