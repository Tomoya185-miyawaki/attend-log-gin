/*
レスポンスの型用のパッケージ（従業員）
*/
package employee

import (
	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/employee"
)

type ListPaginateResponse struct {
	StatusCode  int              `json:"status"`
	CurrentPage int              `json:"currentPage"`
	Emplyees    entity.Employees `json:"employees"`
	LastPage    int              `json:"lastPage"`
}

type SuccessResponse struct {
	StatusCode int              `json:"status"`
	Emplyees   entity.Employees `json:"employees"`
}

type BadListResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
