/*
従業員一覧用のユースケース
*/
package employee

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/employee"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type EmployeeShowUseCase struct{}

func (employeeShowUseCase *EmployeeShowUseCase) Exec(
	c *gin.Context,
) (interface{}, bool) {
	var successResponse response.ShowResponse
	var failedResponse response.BadShowResponse
	// 従業員を取得する
	id := c.Param("id")
	employee, err := repository.NewEmployeeRepository().FetchEmployeeById(id)
	if err != nil {
		log.Warn(err.Error())
		failedResponse.StatusCode = http.StatusBadRequest
		failedResponse.Message = "リクエストが不正です"
		return failedResponse.Failed(), false
	}
	entityEmployee := employee.ConvertToModel()
	successResponse.Emplyee = *entityEmployee
	return successResponse.Success(), true
}
