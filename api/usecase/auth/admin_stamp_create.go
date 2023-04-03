/*
出退勤作成用のユースケース
*/
package auth

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/auth"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/validation"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/auth"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StampCreateUseCase struct{}

func (stampCreateUseCase *StampCreateUseCase) Exec(
	c *gin.Context,
	request *auth.StampCreateRequest,
) *response.CreateStampResponse {
	var err error
	// バリデーションチェック
	if err := validation.ValidationJsonCheck(c, request); err != nil {
		log.Warn(err.Error())
		return &response.CreateStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	// 出退勤を登録する
	err = repository.NewStampRepository().Create(request)
	if err != nil {
		log.Warn(err.Error())
		return &response.CreateStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	return &response.CreateStampResponse{StatusCode: http.StatusOK, Message: "出退勤の登録に成功しました"}
}
