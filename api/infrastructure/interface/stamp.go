/*
Employee用のインターフェース
*/
package IDENT

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
)

type StampInterface interface {
	// 従業員を取得する（ページネーションあり）
	FetchStamps() (*dto.Stamps, error)
}
