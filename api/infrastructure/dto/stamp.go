/*
EmployeeのDTO
*/
package dto

import (
	"time"

	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
)

type Stamp struct {
	ID             uint      `gorm:"primary_key"`
	EmployeeID     uint      `gorm:"index;not null"`
	Status         uint8     `gorm:"not null"`
	StampStartDate time.Time `gorm:"not null"`
	StampEndDate   time.Time
	DeletedAt      time.Time `gorm:"default:null"`
}

type Stamps []Stamp

func (stamp Stamp) ConvertToModel() *entity.Stamp {
	return &entity.Stamp{
		ID:             entity.ID(stamp.ID),
		EmployeeID:     entity.EmployeeID(stamp.EmployeeID),
		Status:         entity.Status(stamp.Status),
		StampStartDate: entity.StampStartDate(helper.TimeToString(stamp.StampStartDate)),
		StampEndDate:   entity.StampEndDate(helper.TimeToString(stamp.StampEndDate)),
	}
}

func (stamps Stamps) ConvertToModel() *entity.Stamps {
	result := make(entity.Stamps, len(stamps))

	for idx, stamp := range stamps {
		stampEntity := entity.Stamp{
			ID:             entity.ID(stamp.ID),
			EmployeeID:     entity.EmployeeID(stamp.EmployeeID),
			Status:         entity.Status(stamp.Status),
			StampStartDate: entity.StampStartDate(helper.TimeToString(stamp.StampStartDate)),
			StampEndDate:   entity.StampEndDate(helper.TimeToString(stamp.StampEndDate)),
		}
		result[idx] = stampEntity
	}

	return &result
}
