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
	// タイムゾーンの設定(日本時間)
	helper.SetAsiaLocation()
	// ルーティングの初期化
	router := gin.Default()
	// corsの設定
	router.Use(middleware.NewCorsConfig())
	// セッションの設定
	router.Use(middleware.InitSession())
	// タイムアウトの設定
	router.Use(middleware.Timeout())
	// ルーティング
	route := router.Group("/api")
	{
		authCtrl := auth.AuthController{}
		employeeCtrl := controller.EmployeeController{}
		stampCtrl := controller.StampController{}
		// ログイン不要なルーティング
		route.POST("/admin/login", authCtrl.Login)                  // ログイン
		route.POST("/admin/password-reset", authCtrl.PasswordReset) // パスワードリセット
		route.GET("/employees", employeeCtrl.List)                  // 従業員一覧取得(ページネーションなし)
		route.GET("/stamps/:employeeId", stampCtrl.Detail)          // 出退勤一覧取得(従業員IDで絞る)
		route.POST("/stamp/create", stampCtrl.Create)               // 出退勤登録(従業員からの登録)
		// ログインが必要なルーティング
		route.Use(middleware.AdminLoginCheck())
		{
			route.POST("/admin/logout", authCtrl.Logout)            // ログアウト
			route.GET("/employee", employeeCtrl.ListPagenate)       // 従業員一覧取得(ページネーション付き)
			route.GET("/employee/:id", employeeCtrl.Show)           // 従業員詳細
			route.POST("/employee/create", employeeCtrl.Create)     // 従業員作成
			route.PATCH("/employee/:id", employeeCtrl.Update)       // 従業員更新
			route.DELETE("/employee/:id", employeeCtrl.Delete)      // 従業員削除
			route.GET("/stamp", stampCtrl.List)                     // 出退勤一覧取得
			route.POST("/admin/stamp/create", authCtrl.StampCreate) // 出退勤登録(管理画面からの登録)
			route.PATCH("/stamp/update", stampCtrl.Update)          // 出退勤更新(管理画面からの更新)
		}
	}
	return router
}
