/*
Admin用のインターフェース
*/
package IDENT

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
)

type AdminInterface interface {
	// メールアドレスを指定して、Adminユーザーを取得
	FindByEmail(email string) (*dto.Admin, error)
	// パスワードを更新する
	UpdatePassword(admin *dto.Admin, password []byte) error
}
