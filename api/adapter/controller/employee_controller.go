package controller

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/employee"
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

func (ec *EmployeeController) Show(c *gin.Context) {
	var employeeShowUseCase usecase.EmployeeShowUseCase
	response, isSuccess := employeeShowUseCase.Exec(c)
	if !isSuccess {
		helper.Response(c, nil, response, http.StatusBadRequest)
	} else {
		helper.Response(c, response, nil, http.StatusOK)
	}
}

func (ec *EmployeeController) Create(c *gin.Context) {
	var request employee.EmployeeCreateRequest
	var employeeCreateUseCase usecase.EmployeeCreateUseCase
	response := employeeCreateUseCase.Exec(c, &request)
	if response.StatusCode != http.StatusOK {
		helper.Response(c, nil, response, response.StatusCode)
	} else {
		helper.Response(c, response, nil, response.StatusCode)
	}
}

func (ec *EmployeeController) Update(c *gin.Context) {
	var request employee.EmployeeUpdateRequest
	var employeeUpdateUseCase usecase.EmployeeUpdateUseCase
	response := employeeUpdateUseCase.Exec(c, &request)
	if response.StatusCode != http.StatusOK {
		helper.Response(c, nil, response, response.StatusCode)
	} else {
		helper.Response(c, response, nil, response.StatusCode)
	}
}
