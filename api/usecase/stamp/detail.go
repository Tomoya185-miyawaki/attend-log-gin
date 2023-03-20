/*
出退勤一覧用のユースケース
*/
package stamp

import (
	"net/http"

	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/stamp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StampDetailUseCase struct{}

func (stampDetailUseCase *StampDetailUseCase) Exec(
	c *gin.Context,
) (interface{}, bool) {
	// クエリパラメータから今日の日付（YYYY-MM-DD）/ページを取得
	today := c.Query("today")
	// 従業員IDを取得する
	employeeId := c.Param("employeeId")

	// 出退勤一覧を取得する
	stampsDto, err := repository.NewStampRepository().FetchStampsByEmployeeId(employeeId, today)
	if err != nil {
		log.Warn(err.Error())
		return badRequestResponse(), false
	}

	// dtoからentityに変換
	stampsEntity := stampsDto.ConvertToModel()
	return successResponse(*stampsEntity), true
}

func badRequestResponse() *response.BadDetailResponse {
	return &response.BadDetailResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "リクエストが不正です",
	}
}

func successResponse(stamps entity.Stamps) *response.DetailResponse {
	return &response.DetailResponse{
		StatusCode: http.StatusOK,
		Stamps:     stamps,
	}
}
