package controller

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
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
