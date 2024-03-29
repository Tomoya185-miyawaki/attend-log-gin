package controller

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/stamp"
	usecase "github.com/Tomoya185-miyawaki/attend-log-gin/usecase/stamp"
	"github.com/gin-gonic/gin"
)

type StampController struct{}

func (sc *StampController) List(c *gin.Context) {
	var stampListUseCase usecase.StampListUseCase
	response, isSuccess := stampListUseCase.Exec(c)
	if !isSuccess {
		helper.Response(c, nil, response, http.StatusBadRequest)
	} else {
		helper.Response(c, response, nil, http.StatusOK)
	}
}

func (sc *StampController) Detail(c *gin.Context) {
	var stampListUseCase usecase.StampDetailUseCase
	response, isSuccess := stampListUseCase.Exec(c)
	if !isSuccess {
		helper.Response(c, nil, response, http.StatusBadRequest)
	} else {
		helper.Response(c, response, nil, http.StatusOK)
	}
}

func (sc *StampController) Create(c *gin.Context) {
	var request stamp.StampCreateRequest
	var stampCreateUseCase usecase.StampCreateUseCase
	response := stampCreateUseCase.Exec(c, &request)
	if response.StatusCode != http.StatusOK {
		helper.Response(c, nil, response, response.StatusCode)
	} else {
		helper.Response(c, response, nil, response.StatusCode)
	}
}

func (sc *StampController) Update(c *gin.Context) {
	var request stamp.StampUpdateRequest
	var stampUpdateUseCase usecase.StampUpdateUseCase
	response := stampUpdateUseCase.Exec(c, &request)
	if response.StatusCode != http.StatusOK {
		helper.Response(c, nil, response, response.StatusCode)
	} else {
		helper.Response(c, response, nil, response.StatusCode)
	}
}

func (sc *StampController) Delete(c *gin.Context) {
	var stampDeleteUseCase usecase.StampDeleteUseCase
	response := stampDeleteUseCase.Exec(c)
	if response.StatusCode != http.StatusOK {
		helper.Response(c, nil, response, response.StatusCode)
	} else {
		helper.Response(c, response, nil, response.StatusCode)
	}
}
