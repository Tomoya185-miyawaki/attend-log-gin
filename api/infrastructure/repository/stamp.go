/*
employeesテーブルを扱うための実態
*/
package repository

import (
	"errors"

	db "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	IDENT "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/interface"
	log "github.com/sirupsen/logrus"
)

type stampRepository struct{}

func NewStampRepository() IDENT.StampInterface {
	return &stampRepository{}
}

func (sr *stampRepository) FetchStamps() (*dto.Stamps, error) {
	db := db.GetDB()
	// stamp := &dto.Stamp{}
	stamps := &dto.Stamps{}

	if err := db.Find(&stamps).Error; err != nil {
		log.Warn("aaa")
		log.Warn(err.Error())
		return nil, errors.New("出退勤が取得できませんでした")
	}
	log.Warn("bb")
	log.Warn(*stamps)
	return stamps, nil
}
