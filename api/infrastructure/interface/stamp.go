/*
Employee用のインターフェース
*/
package IDENT

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/auth"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/stamp"
)

type StampInterface interface {
	// 出退勤情報を取得する
	FetchStamps(today string) (*dto.Stamps, error)
	// 従業員IDに紐づく出退勤情報を取得する
	FetchStampsByEmployeeId(employeeId string, today string) (*dto.Stamps, error)
	// カウントを取得する（ページネーション用）
	FetchCount(today string) (*dto.StampCount, error)
	// 出勤・休憩開始を登録する
	Create(request *auth.StampCreateRequest) error
	// 出勤・休憩開始を登録する
	CreateStampStart(request *stamp.StampCreateRequest) error
	// 出退勤を更新する
	Update(request *stamp.StampUpdateRequest) error
	// 退勤・休憩終了を登録する
	UpdateStampEnd(request *stamp.StampCreateRequest, checkStatus int) error
	// 出退勤を削除する
	Delete(id string) error
}
