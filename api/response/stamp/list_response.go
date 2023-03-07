/*
レスポンスの型用のパッケージ（出退勤）
*/
package stamp

import (
	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
)

type ListResponse struct {
	CurrentPage int                         `json:"currentPage"`
	Stamps      map[string]entity.StampDate `json:"stamps"`
	EmployeeIds []uint                      `json:"employeeIds"`
	LastPage    int                         `json:"lastPage"`
}

type BadListResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func (listResponse *ListResponse) Success() *ListResponse {
	return listResponse
}

func (badListResponse *BadListResponse) Failed() *BadListResponse {
	return badListResponse
}
