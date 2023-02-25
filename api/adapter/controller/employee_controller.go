package controller

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	usecase "github.com/Tomoya185-miyawaki/attend-log-gin/usecase/employee"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct{}

func (ec *EmployeeController) List(c *gin.Context) {
	var employeeListUseCase usecase.EmployeeListUseCase
	response, isSuccess := employeeListUseCase.Exec(c)
	if !isSuccess {
		helper.Response(c, nil, response, http.StatusBadRequest)
	} else {
		helper.Response(c, response, nil, http.StatusOK)
	}
}
