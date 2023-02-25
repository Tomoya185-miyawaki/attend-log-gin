/*
ルーティング用パッケージ
*/
package router

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/adapter/controller"
	"github.com/Tomoya185-miyawaki/attend-log-gin/adapter/controller/auth"
	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/middleware"
	"github.com/gin-gonic/gin"
)

// ルーティング処理を行う
func Bind() *gin.Engine {
	// ログの設定
	helper.SetUpLog()
	// ルーティングの初期化
	router := gin.Default()
	// corsの設定
	router.Use(middleware.NewCorsConfig())
	// セッションの設定
	router.Use(middleware.InitSession())
	// ルーティング
	route := router.Group("/api")
	{
		authCtrl := auth.AuthController{}
		employeeCtrl := controller.EmployeeController{}
		// ログイン不要なルーティング
		route.POST("/admin/login", authCtrl.Login)
		// ログインが必要なルーティング
		route.Use(middleware.AdminLoginCheck())
		{
			route.POST("/admin/logout", authCtrl.Logout)
			route.GET("/employee", employeeCtrl.List)
		}
	}
	return router
}
