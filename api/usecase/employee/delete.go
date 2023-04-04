/*
従業員削除用のユースケース
*/
package employee

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/employee"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type EmployeeDeleteUseCase struct{}

func (employeeDeleteUseCase *EmployeeDeleteUseCase) Exec(
	c *gin.Context,
) *response.DeleteEmployeeResponse {
	// 従業員を削除する
	id := c.Param("id")
	if err := repository.NewEmployeeRepository().Delete(id); err != nil {
		log.Warn(err.Error())
		return &response.DeleteEmployeeResponse{StatusCode: http.StatusBadRequest, Message: "リクエストが不正です"}
	}
	return &response.DeleteEmployeeResponse{StatusCode: http.StatusOK, Message: "従業員の削除に成功しました"}
}
