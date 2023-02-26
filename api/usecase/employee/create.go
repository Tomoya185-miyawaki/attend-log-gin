/*
従業員作成用のユースケース
*/
package employee

import (
	"net/http"
	"strconv"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/employee"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/validation"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/employee"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type EmployeeCreateUseCase struct{}

func (employeeCreateUseCase *EmployeeCreateUseCase) Exec(
	c *gin.Context,
	request *employee.EmployeeCreateRequest,
) *response.CreateEmployeeResponse {
	// バリデーションチェック
	isOk, err := validation.ValidationCheck(c, request)
	if !isOk {
		log.Warn(err.Error())
		return &response.CreateEmployeeResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	// 時給の型変換
	hourlyWage, err := strconv.Atoi(request.HourlyWage)
	if err != nil {
		log.Warn(err.Error())
		return &response.CreateEmployeeResponse{StatusCode: http.StatusBadRequest, Message: "時給の値が不正です"}
	}
	// 従業員を作成する
	if err := repository.NewEmployeeRepository().Create(request, uint(hourlyWage)); err != nil {
		log.Warn(err.Error())
		return &response.CreateEmployeeResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	return &response.CreateEmployeeResponse{StatusCode: http.StatusOK, Message: "従業員の作成に成功しました"}
}
