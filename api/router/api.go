/*
ルーティング用パッケージ
*/
package router

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/adapter/controller/auth"
	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	cors "github.com/Tomoya185-miyawaki/attend-log-gin/middleware"
	session "github.com/Tomoya185-miyawaki/attend-log-gin/middleware"
	"github.com/gin-gonic/gin"
)

// ルーティング処理を行う
func Bind() *gin.Engine {
	// ログの設定
	helper.SetUpLog()
	// ルーティングの初期化
	router := gin.Default()
	// corsの設定
	router.Use(cors.NewCorsConfig())
	// セッションの設定
	router.Use(session.InitSession())
	// ルーティング
	route := router.Group("/api")
	{
		// ログイン不要なルーティング
		authCtrl := auth.AuthController{}
		route.POST("/admin/login", authCtrl.Login)
	}
	return router
}
