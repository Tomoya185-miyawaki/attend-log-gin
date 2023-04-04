/*
出退勤削除用のユースケース
*/
package stamp

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/stamp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StampDeleteUseCase struct{}

func (stampDeleteUseCase *StampDeleteUseCase) Exec(
	c *gin.Context,
) *response.DeleteStampResponse {
	// 出退勤を削除する
	id := c.Param("id")
	if err := repository.NewStampRepository().Delete(id); err != nil {
		log.Warn(err.Error())
		return &response.DeleteStampResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	return &response.DeleteStampResponse{StatusCode: http.StatusOK, Message: "出退勤の削除に成功しました"}
}
