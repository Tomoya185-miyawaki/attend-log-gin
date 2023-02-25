/*
ログイン用のユースケース
*/
package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/auth"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginUseCase struct{}

func (adminLoginUseCase *AdminLoginUseCase) Exec(
	c *gin.Context,
	request *auth.LoginRequest,
) *response.LoginResponse {
	badRequestMessage := "リクエストが不正です"
	// バリデーションチェック
	isOk := validationCheck(c, request)
	if !isOk {
		return &response.LoginResponse{StatusCode: http.StatusBadRequest, Message: badRequestMessage}
	}
	// リクエストパラメータのEmailに紐づくユーザーが存在するかどうか
	admin, err := repository.NewAdminRepository().FindByEmail(request.Email)
	if err != nil {
		log.Warn(err.Error())
		return &response.LoginResponse{StatusCode: http.StatusBadRequest, Message: badRequestMessage}
	}
	// リクエストパラメータのPasswordが正しいかどうか
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password))
	if err != nil {
		log.Warn(err.Error())
		return &response.LoginResponse{StatusCode: http.StatusBadRequest, Message: badRequestMessage}
	}
	// セッションに情報を保存できるかどうか
	session := sessions.Default(c)
	loginAdminUser, err := json.Marshal(admin)
	if err != nil {
		log.Warn(err.Error())
		return &response.LoginResponse{StatusCode: http.StatusInternalServerError, Message: "サーバー側でエラーが発生しました"}
	}
	// セッションにログインユーザを保存
	session.Set("loginAdminUser", string(loginAdminUser))
	session.Save()
	return &response.LoginResponse{StatusCode: http.StatusOK, Message: "ログインに成功しました"}
}

func validationCheck(
	c *gin.Context,
	request *auth.LoginRequest,
) bool {
	// バリデーションチェック
	if err := c.BindJSON(&request); err != nil {
		log.Warn(err.Error())
		return false
	}
	return true
}
