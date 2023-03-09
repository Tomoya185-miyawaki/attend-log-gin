/*
出退勤一覧用のユースケース
*/
package stamp

import (
	"math"
	"net/http"
	"strconv"
	"time"

	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/stamp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StampListUseCase struct{}

func (stampListUseCase *StampListUseCase) Exec(
	c *gin.Context,
) (interface{}, bool) {
	var successResponse response.ListResponse
	var failedResponse response.BadListResponse
	// クエリパラメータから今日の日付（YYYY-MM-DD）/ページを取得
	today := c.Query("today")
	page := c.Query("page")
	if page == "" {
		// 空の場合は1ページ目を設定
		page = "1"
	}

	// responseのために型変換をする
	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Warn(err.Error())
		failedResponse.StatusCode = http.StatusBadRequest
		failedResponse.Message = "pageの型変換に失敗しました"
		return failedResponse.Failed(), false
	}

	// 出退勤一覧を取得する
	stamps, err := repository.NewStampRepository().FetchStamps(today)
	if err != nil {
		log.Warn(err.Error())
		failedResponse.StatusCode = http.StatusBadRequest
		failedResponse.Message = err.Error()
		return failedResponse.Failed(), false
	}

	// 出退勤の件数を取得する（ページネーション用）
	count, err := repository.NewStampRepository().FetchCount(today)
	if err != nil {
		log.Warn(err.Error())
		failedResponse.StatusCode = http.StatusBadRequest
		failedResponse.Message = err.Error()
		return failedResponse.Failed(), false
	}

	// トータルのページを設定
	limit := 10
	totalPages := int(math.Ceil(float64(count.Cnt) / float64(limit)))
	// dtoからentityに変換
	stampsEntity := stamps.ConvertToModel()
	// 従業員IDを取得
	employeeIdMap := map[entity.EmployeeID]entity.EmployeeID{}
	for _, stamp := range *stampsEntity {
		employeeIdMap[stamp.EmployeeID] = stamp.EmployeeID
	}
	// 従業員を取得する
	var employeeIds []uint
	for _, employeeId := range employeeIdMap {
		employeeIds = append(employeeIds, uint(employeeId))
	}
	employees, err := repository.NewEmployeeRepository().FetchEmployeeByIds(employeeIds)
	if err != nil {
		log.Warn(err.Error())
		failedResponse.StatusCode = http.StatusBadRequest
		failedResponse.Message = err.Error()
		return failedResponse.Failed(), false
	}
	// 出勤・退勤・休憩・労働時間に直す
	var employeeNameId map[uint]string = make(map[uint]string)
	for _, employee := range *employees {
		employeeNameId[employee.ID] = employee.Name
	}

	resStampsDate := getAttendRestDate(stamps, employeeNameId)

	// レスポンス設定
	successResponse.Stamps = resStampsDate
	successResponse.CurrentPage = intPage
	successResponse.EmployeeIds = employeeIds
	successResponse.LastPage = totalPages
	return successResponse.Success(), true
}

func getAttendRestDate(stamps *dto.Stamps, employeeNameId map[uint]string) map[string]entity.StampDate {
	stampDate := entity.NewStampDate()
	var attendDate time.Time    // 出勤時刻
	var leavingDate *time.Time  // 退勤時刻
	var restStartDate time.Time //休憩開始時刻
	var restEndDate *time.Time  //休憩終了時刻
	var resStampsDate map[string]entity.StampDate = make(map[string]entity.StampDate)
	for _, stamp := range *stamps {
		if _, ok := employeeNameId[uint(stamp.EmployeeID)]; !ok {
			continue
		}
		if stamp.Status == dto.Attend {
			stampDate.SetAttendDate(entity.AttendDate(helper.TimeToStringDuration(stamp.StampStartDate)))
			attendDate = stamp.StampStartDate
			if stamp.StampEndDate != nil {
				stampDate.SetLeavingDate(entity.LeavingDate(helper.TimeToStringDurationPointer(stamp.StampEndDate)))
				leavingDate = stamp.StampEndDate
			} else {
				leavingDate = nil
			}
		}
		if stamp.Status == dto.Rest {
			restStartDate = stamp.StampStartDate
			restEndDate = stamp.StampEndDate
			if !restEndDate.IsZero() {
				restDate := restEndDate.Sub(restStartDate)
				if restDate.Hours() >= 1 {
					stampDate.SetRestDate(entity.RestDate(strconv.Itoa(int(restDate.Hours())) + "時間" + strconv.Itoa(int(restDate.Minutes()-(restDate.Hours()*60))) + "分"))
				} else {
					stampDate.SetRestDate(entity.RestDate(strconv.Itoa(int(restDate.Minutes())) + "分"))
				}
			} else {
				stampDate.SetRestDate("-")
			}
			stampDate.SetWorkingDate(entity.WorkingDate(helper.TimeToStringDurationPointer(stamp.StampEndDate)))
		}
		if leavingDate != nil {
			var workingMinutes float64 // 労働時間（分）
			workingDate := leavingDate.Sub(attendDate)
			if !restEndDate.IsZero() {
				restDate := restEndDate.Sub(restStartDate)
				workingMinutes = workingDate.Minutes() - restDate.Minutes()
			} else {
				workingMinutes = workingDate.Minutes()
			}
			if int(workingMinutes)/60 >= 1 {
				stampDate.SetWorkingDate(entity.WorkingDate(strconv.Itoa(int(workingMinutes)/60)+"時間"+strconv.Itoa(int(workingMinutes)%60)) + "分")
			} else {
				stampDate.SetWorkingDate(entity.WorkingDate(strconv.Itoa(int(workingMinutes)) + "分"))
			}
		}
		resStampsDate[employeeNameId[uint(stamp.EmployeeID)]] = *stampDate
	}
	return resStampsDate
}
