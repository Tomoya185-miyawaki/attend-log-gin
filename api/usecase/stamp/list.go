/*
出退勤一覧用のユースケース
*/
package stamp

import (
	"math"
	"net/http"
	"strconv"

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

	stampDate := entity.NewStampDate()
	var resStampsDate map[string]entity.StampDate = make(map[string]entity.StampDate)
	for _, stamp := range *stamps {
		if _, ok := employeeNameId[uint(stamp.EmployeeID)]; !ok {
			continue
		}
		if stamp.Status == dto.Attend {
			stampDate.SetAttendDate(entity.AttendDate(helper.TimeToStringDuration(stamp.StampStartDate)))
			stampDate.SetLeavingDate(entity.LeavingDate(helper.TimeToStringDuration(stamp.StampEndDate)))
		}
		if stamp.Status == dto.Rest {
			restStartDate := stamp.StampStartDate
			restEndDate := stamp.StampEndDate
			sub := restEndDate.Sub(restStartDate)
			if !restEndDate.IsZero() {
				stampDate.SetRestDate(entity.RestDate(strconv.Itoa(int(sub.Minutes())) + "分"))
			} else {
				stampDate.SetRestDate("-")
			}
			stampDate.SetWorkingDate(entity.WorkingDate(helper.TimeToStringDuration(stamp.StampEndDate)))
		}
		resStampsDate[employeeNameId[uint(stamp.EmployeeID)]] = *stampDate
	}

	// レスポンス設定
	successResponse.Stamps = resStampsDate
	successResponse.CurrentPage = intPage
	successResponse.EmployeeIds = employeeIds
	successResponse.LastPage = totalPages
	return successResponse.Success(), true
}
