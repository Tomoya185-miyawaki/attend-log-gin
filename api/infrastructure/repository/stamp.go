/*
employeesテーブルを扱うための実態
*/
package repository

import (
	"errors"
	"time"

	db "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	IDENT "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/interface"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/auth"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/stamp"
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

func (sr *stampRepository) FetchStampsByEmployeeId(employeeId string, today string) (*dto.Stamps, error) {
	db := db.GetDB()
	stamps := &dto.Stamps{}

	if err := db.Where("employee_id = ?", employeeId).Where("stamp_start_date like ?", "%"+today+"%").Find(&stamps).Error; err != nil {
		log.Warn(err.Error())
		return nil, errors.New("従業員IDに紐づく出退勤が取得できませんでした")
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

func (sr *stampRepository) Create(request *auth.StampCreateRequest) error {
	db := db.GetDB()
	stamp := &dto.Stamp{
		EmployeeID:     uint(request.EmployeeId),
		Status:         dto.StampStatus(request.Status),
		StampStartDate: request.StampStartDate,
		StampEndDate:   request.StampEndDate,
	}
	if err := db.Create(&stamp).Error; err != nil {
		return errors.New("出退勤を作成できませんでした")
	}
	return nil
}

func (sr *stampRepository) CreateStampStart(request *stamp.StampCreateRequest) error {
	db := db.GetDB()
	stamp := &dto.Stamp{
		EmployeeID:     uint(request.EmployeeId),
		Status:         dto.StampStatus(request.Status),
		StampStartDate: time.Now(),
	}
	if err := db.Create(&stamp).Error; err != nil {
		return errors.New("出退勤を作成できませんでした")
	}
	return nil
}

func (sr *stampRepository) Update(request *stamp.StampUpdateRequest) error {
	db := db.GetDB()
	stamp := &dto.Stamp{}

	if err := db.Where("id = ?", request.Id).Find(&stamp).Error; err != nil {
		log.Warn(err.Error())
		return errors.New("出勤もしくは休憩開始が取得できませんでした")
	}

	stamp.StampStartDate = request.StampStartDate
	stamp.StampEndDate = request.StampEndDate
	if err := db.Save(&stamp).Error; err != nil {
		return errors.New("出退勤を更新できませんでした")
	}
	return nil
}

func (sr *stampRepository) UpdateStampEnd(request *stamp.StampCreateRequest, checkStatus int) error {
	db := db.GetDB()
	stamp := &dto.Stamp{}

	if err := db.Where("employee_id = ?", request.EmployeeId).Where("status = ?", checkStatus).Where("stamp_start_date like ?", "%"+request.Today+"%").Find(&stamp).Error; err != nil {
		log.Warn(err.Error())
		return errors.New("出勤もしくは休憩開始が取得できませんでした")
	}

	stamp.StampEndDate = time.Now()
	if err := db.Save(&stamp).Error; err != nil {
		return errors.New("従業員を更新できませんでした")
	}
	return nil
}

func (sr *stampRepository) Delete(id string) error {
	db := db.GetDB()
	stamp := &dto.Stamp{}

	if err := db.Delete(&stamp, id).Error; err != nil {
		return errors.New("出退勤の削除に失敗しました")
	}
	return nil
}
