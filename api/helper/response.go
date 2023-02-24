/*
レスポンス用ヘルパー
*/
package helper

import (
	"github.com/gin-gonic/gin"
)

func Response(
	c *gin.Context,
	result interface{},
	err interface{},
	statusCode int,
) {
	if err != nil {
		c.JSON(statusCode, err)
	} else {
		c.JSON(statusCode, result)
	}
}
