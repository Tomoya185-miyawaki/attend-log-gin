/*
レスポンスの型用のパッケージ（出退勤）
*/
package stamp

import (
	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
)

type ListResponse struct {
	CurrentPage int           `json:"currentPage"`
	Stamps      entity.Stamps `json:"stamps"`
	LastPage    int           `json:"lastPage"`
}

type BadListResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
