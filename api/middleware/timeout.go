/*
APIタイムアウト用のミドルウェア
*/
package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/timeout"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func Timeout() gin.HandlerFunc {
	// env読み込み
	helper.GetEnv()
	apiTimeout, err := strconv.Atoi(os.Getenv("API_TIMEOUT"))
	if err != nil {
		log.Warn(err.Error())
		apiTimeout = 5
	}
	return timeout.New(
		timeout.WithTimeout(time.Duration(apiTimeout)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	c.JSON(
		http.StatusRequestTimeout,
		&response.TimeoutResponse{StatusCode: http.StatusRequestTimeout, Message: "タイムアウトエラー"},
	)
}
