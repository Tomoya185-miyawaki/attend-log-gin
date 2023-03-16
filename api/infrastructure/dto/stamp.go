/*
EmployeeのDTO
*/
package dto

import (
	"time"

	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"gorm.io/gorm"
)

type Stamp struct {
	ID             uint        `gorm:"primary_key"`
	EmployeeID     uint        `gorm:"index;not null"`
	Status         StampStatus `gorm:"not null"`
	StampStartDate time.Time   `gorm:"not null"`
	StampEndDate   time.Time
	DeletedAt      gorm.DeletedAt `gorm:"default:null"`
}

type Stamps []Stamp

type StampCount struct {
	Cnt int
}

type StampStatus uint8

const (
	Attend = StampStatus(1)
	Rest   = StampStatus(2)
)

type StampRequestStatus int

const (
	AttendStart = StampRequestStatus(1)
	RestStart   = StampRequestStatus(2)
	RestEnd     = StampRequestStatus(3)
	WorkingEnd  = StampRequestStatus(4)
)

func (stamp Stamp) ConvertToModel() *entity.Stamp {
	if !stamp.StampEndDate.IsZero() {
		return &entity.Stamp{
			ID:             entity.ID(stamp.ID),
			EmployeeID:     entity.EmployeeID(stamp.EmployeeID),
			Status:         entity.Status(stamp.Status),
			StampStartDate: entity.StampStartDate(helper.TimeToString(stamp.StampStartDate)),
			StampStart:     entity.StampStart(helper.TimeToStringDuration(stamp.StampStartDate)),
			StampEndDate:   entity.StampEndDate(helper.TimeToString(stamp.StampEndDate)),
			StampEnd:       entity.StampEnd(helper.TimeToStringDuration(stamp.StampEndDate)),
		}
	}
	return &entity.Stamp{
		ID:             entity.ID(stamp.ID),
		EmployeeID:     entity.EmployeeID(stamp.EmployeeID),
		Status:         entity.Status(stamp.Status),
		StampStartDate: entity.StampStartDate(helper.TimeToString(stamp.StampStartDate)),
		StampStart:     entity.StampStart(helper.TimeToStringDuration(stamp.StampStartDate)),
	}
}

func (stamps Stamps) ConvertToModel() *entity.Stamps {
	result := make(entity.Stamps, len(stamps))

	for idx, stamp := range stamps {
		var stampEntity entity.Stamp
		if !stamp.StampEndDate.IsZero() {
			stampEntity = entity.Stamp{
				ID:             entity.ID(stamp.ID),
				EmployeeID:     entity.EmployeeID(stamp.EmployeeID),
				Status:         entity.Status(stamp.Status),
				StampStartDate: entity.StampStartDate(helper.TimeToString(stamp.StampStartDate)),
				StampStart:     entity.StampStart(helper.TimeToStringDuration(stamp.StampStartDate)),
				StampEndDate:   entity.StampEndDate(helper.TimeToString(stamp.StampEndDate)),
				StampEnd:       entity.StampEnd(helper.TimeToString(stamp.StampEndDate)),
			}
		} else {
			stampEntity = entity.Stamp{
				ID:             entity.ID(stamp.ID),
				EmployeeID:     entity.EmployeeID(stamp.EmployeeID),
				Status:         entity.Status(stamp.Status),
				StampStartDate: entity.StampStartDate(helper.TimeToString(stamp.StampStartDate)),
				StampStart:     entity.StampStart(helper.TimeToStringDuration(stamp.StampStartDate)),
			}
		}
		result[idx] = stampEntity
	}

	return &result
}
