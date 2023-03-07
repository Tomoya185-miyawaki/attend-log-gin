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

func (sr *stampRepository) FetchStamps(today string) (*dto.Stamps, error) {
	db := db.GetDB()
	stamps := &dto.Stamps{}

	if err := db.Where("stamp_start_date like ?", "%"+today+"%").Find(&stamps).Error; err != nil {
		log.Warn(err.Error())
		return nil, errors.New("出退勤が取得できませんでした")
	}
	return stamps, nil
}

func (sr *stampRepository) FetchCount(today string) (*dto.StampCount, error) {
	db := db.GetDB()
	stampCount := &dto.StampCount{}

	if err := db.Table("stamps").Select("distinct employee_id as cnt").Where("status = ?", dto.Attend).Where("stamp_start_date like ?", "%"+today+"%").Scan(&stampCount).Error; err != nil {
		if stampCount.Cnt == 0 {
			return stampCount, nil
		}
		log.Warn(err.Error())
		return nil, errors.New("出退勤カウントが取得できませんでした")
	}
	return stampCount, nil
}
