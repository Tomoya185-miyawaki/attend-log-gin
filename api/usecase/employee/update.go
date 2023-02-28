/*
従業員更新用のユースケース
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

type EmployeeUpdateUseCase struct{}

func (employeeUpdateUseCase *EmployeeUpdateUseCase) Exec(
	c *gin.Context,
	request *employee.EmployeeUpdateRequest,
) *response.UpdateEmployeeResponse {
	// バリデーションチェック
	if err := validation.ValidationJsonCheck(c, request); err != nil {
		log.Warn(err.Error())
		return &response.UpdateEmployeeResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	// 時給の型変換
	hourlyWage, err := strconv.Atoi(request.HourlyWage)
	if err != nil {
		log.Warn(err.Error())
		return &response.UpdateEmployeeResponse{StatusCode: http.StatusBadRequest, Message: "時給の値が不正です"}
	}
	id := c.Param("id")
	// 従業員を更新する
	if err := repository.NewEmployeeRepository().Update(id, request.Name, uint(hourlyWage)); err != nil {
		log.Warn(err.Error())
		return &response.UpdateEmployeeResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	return &response.UpdateEmployeeResponse{StatusCode: http.StatusOK, Message: "従業員の更新に成功しました"}
}
