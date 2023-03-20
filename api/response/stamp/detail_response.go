/*
レスポンスの型用のパッケージ（出退勤）
*/
package stamp

import (
	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
)

type DetailResponse struct {
	StatusCode int           `json:"status"`
	Stamps     entity.Stamps `json:"stamps"`
}

type BadDetailResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
