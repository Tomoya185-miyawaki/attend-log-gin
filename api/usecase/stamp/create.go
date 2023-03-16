/*
出退勤一覧用のユースケース
*/
package stamp

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/stamp"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/validation"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/stamp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StampCreateUseCase struct{}

func (stampCreateUseCase *StampCreateUseCase) Exec(
	c *gin.Context,
	request *stamp.StampCreateRequest,
) *response.CreateStampResponse {
	var err error
	// バリデーションチェック
	if err := validation.ValidationJsonCheck(c, request); err != nil {
		log.Warn(err.Error())
		return &response.CreateStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	// 出退勤を登録する
	requestStatus := dto.StampRequestStatus(request.Status)
	if requestStatus == dto.AttendStart || requestStatus == dto.RestStart {
		err = repository.NewStampRepository().Create(request)
	} else if requestStatus == dto.WorkingEnd || requestStatus == dto.RestEnd {
		var checkStatus int
		if requestStatus == dto.WorkingEnd {
			checkStatus = int(dto.Attend)
		} else {
			checkStatus = int(dto.Rest)
		}
		err = repository.NewStampRepository().Update(request, checkStatus)
	} else {
		log.Warn("statusが不正です")
		return &response.CreateStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	if err != nil {
		log.Warn(err.Error())
		return &response.CreateStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	return &response.CreateStampResponse{StatusCode: http.StatusOK, Message: "出退勤の登録に成功しました"}
}
