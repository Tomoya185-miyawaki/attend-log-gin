/*
Employee用のインターフェース
*/
package IDENT

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/stamp"
)

type StampInterface interface {
	// 出退勤情報を取得する
	FetchStamps(today string) (*dto.Stamps, error)
	// カウントを取得する（ページネーション用）
	FetchCount(today string) (*dto.StampCount, error)
	// 出退勤を登録する
	Create(request *stamp.StampCreateRequest) error
}
