/*
出退勤一覧用のユースケース
*/
package stamp

import (
	"net/http"

	entity "github.com/Tomoya185-miyawaki/attend-log-gin/entity/stamp"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/repository"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/stamp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StampListUseCase struct{}

func (stampListUseCase *StampListUseCase) Exec(
	c *gin.Context,
) (interface{}, bool) {
	// 出退勤一覧を取得する
	stamps, err := repository.NewStampRepository().FetchStamps()
	if err != nil {
		log.Warn(err.Error())
		return badRequestResponse(), false
	}
	// totalPages := int(math.Ceil(float64(total) / float64(limit)))
	stampEmployees := stamps.ConvertToModel()
	return successRequestResponse(1, *stampEmployees, 1), true
}

func badRequestResponse() *response.BadListResponse {
	return &response.BadListResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "リクエストが不正です",
	}
}

func successRequestResponse(
	currentPage int,
	stamps entity.Stamps,
	lastPage int,
) *response.ListResponse {
	return &response.ListResponse{
		CurrentPage: currentPage,
		Stamps:      stamps,
		LastPage:    lastPage,
	}
}
