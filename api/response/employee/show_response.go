/*
レスポンスの型用のパッケージ（従業員詳細）
*/
package employee

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/entity"
)

type ShowResponse struct {
	Emplyee entity.Employee `json:"employee"`
}

type BadShowResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func (showResponse *ShowResponse) Success() *ShowResponse {
	return showResponse
}

func (badShowResponse *BadShowResponse) Failed() *BadShowResponse {
	return badShowResponse
}
