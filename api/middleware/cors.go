/*
CORS用ミドルウェア
*/
package middleware

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 許可したいオリジン
		AllowOrigins: []string{
			os.Getenv("FRONT_DOMAIN"),
		},
		// 許可したいHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	})
}
