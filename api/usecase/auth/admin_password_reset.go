/*
パスワードリセット用のユースケース
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
	"golang.org/x/crypto/bcrypt"
)

type AdminPasswordResetUseCase struct{}

func (adminPasswordResetUseCase *AdminPasswordResetUseCase) Exec(
	c *gin.Context,
	request *auth.PasswordResetRequest,
) *response.PasswordResetResponse {
	badRequestMessage := "リクエストが不正です"
	// バリデーションチェック
	isOk, err := validation.ValidationCheck(c, request)
	if !isOk {
		log.Warn(err.Error())
		return &response.PasswordResetResponse{StatusCode: http.StatusBadRequest, Message: badRequestMessage}
	}
	// リクエストパラメータのEmailに紐づくユーザーが存在するかどうか
	admin, err := repository.NewAdminRepository().FindByEmail(request.Email)
	if err != nil {
		log.Warn(err.Error())
		return &response.PasswordResetResponse{StatusCode: http.StatusBadRequest, Message: badRequestMessage}
	}
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		log.Warn(err.Error())
		return &response.PasswordResetResponse{StatusCode: http.StatusBadRequest, Message: badRequestMessage}
	}
	// パスワードを更新する
	if err := repository.NewAdminRepository().UpdatePassword(admin, hashedPassword); err != nil {
		log.Warn(err.Error())
		return &response.PasswordResetResponse{StatusCode: http.StatusBadRequest, Message: err.Error()}
	}
	return &response.PasswordResetResponse{StatusCode: http.StatusOK, Message: "パスワードリセットに成功しました"}
}
