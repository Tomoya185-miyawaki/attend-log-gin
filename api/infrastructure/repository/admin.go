/*
Adminsテーブルを扱うための実態
*/
package repository

import (
	"errors"

	db "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	IDENT "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/interface"
)

type adminRepository struct{}

func NewAdminRepository() IDENT.AdminInterface {
	return &adminRepository{}
}

func (ar *adminRepository) FindByEmail(email string) (*dto.Admin, error) {
	db := db.GetDB()
	admin := &dto.Admin{}

	if db.Where("email = ?", email).First(admin).RecordNotFound() {
		return nil, errors.New("ユーザーが見つかりませんでした")
	}
	return admin, nil
}

func (ar *adminRepository) UpdatePassword(admin *dto.Admin, password []byte) error {
	db := db.GetDB()

	if err := db.Model(admin).Update("password", password).Error; err != nil {
		return errors.New("パスワードの更新に失敗しました")
	}
	return nil
}
