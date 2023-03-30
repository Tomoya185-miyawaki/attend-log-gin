/*
出退勤更新用のユースケース
*/
package stamp

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/stamp"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/validation"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/stamp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StampUpdateUseCase struct{}

func (stampUpdateUseCase *StampUpdateUseCase) Exec(
	c *gin.Context,
	request *stamp.StampUpdateRequest,
) *response.UpdateStampResponse {
	var err error
	// バリデーションチェック
	if err := validation.ValidationJsonCheck(c, request); err != nil {
		log.Warn(err.Error())
		return &response.UpdateStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	// 出退勤を更新する
	err = repository.NewStampRepository().Update(request)
	if err != nil {
		log.Warn(err.Error())
		return &response.UpdateStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	return &response.UpdateStampResponse{StatusCode: http.StatusOK, Message: "出退勤の更新に成功しました"}
}
