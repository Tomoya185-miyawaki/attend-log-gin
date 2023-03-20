/*
従業員一覧用のユースケース
*/
package employee

import (
	"math"
	"net/http"
	"strconv"

	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/employee"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/employee"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type EmployeeListUseCase struct{}

func (employeeListUseCase *EmployeeListUseCase) Exec(
	c *gin.Context,
) (interface{}, bool) {
	// 従業員一覧を取得する
	employees, err := repository.NewEmployeeRepository().FetchEmployees()
	if err != nil {
		log.Warn(err.Error())
		return badRequestResponse(), false
	}
	entityEmployees := employees.ConvertToModel()
	return successResponse(*entityEmployees), true
}

func (employeeListUseCase *EmployeeListUseCase) ExecPaginate(
	c *gin.Context,
) (interface{}, bool) {
	// クエリパラメータからpageを取得
	page := c.Query("page")
	if page == "" {
		// 空の場合は1ページ目を設定
		page = "1"
	}
	// responseのために型変換をする
	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Warn(err.Error())
		return badRequestResponse(), false
	}
	limit := 10
	offset := limit * (intPage - 1)
	if intPage == 1 {
		offset = -1
	}
	// 従業員一覧を取得する
	employees, err, total := repository.NewEmployeeRepository().FetchEmployeesByPaginate(intPage, limit, offset)
	if err != nil {
		log.Warn(err.Error())
		return badRequestResponse(), false
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	entityEmployees := employees.ConvertToModel()
	return successPaginateResponse(intPage, *entityEmployees, totalPages), true
}

func badRequestResponse() *response.BadListResponse {
	return &response.BadListResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "リクエストが不正です",
	}
}

func successResponse(employees entity.Employees) *response.SuccessResponse {
	return &response.SuccessResponse{
		StatusCode: http.StatusOK,
		Emplyees:   employees,
	}
}

func successPaginateResponse(
	currentPage int,
	employees entity.Employees,
	lastPage int,
) *response.ListPaginateResponse {
	return &response.ListPaginateResponse{
		StatusCode:  http.StatusOK,
		CurrentPage: currentPage,
		Emplyees:    employees,
		LastPage:    lastPage,
	}
}
